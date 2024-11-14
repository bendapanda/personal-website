package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

const (
	cacheExparation = 5 * time.Minute
	cachePurge      = 10 * time.Minute
)

var db *sql.DB
var c = cache.New(cacheExparation, cachePurge)
var commentCount int
var timeSinceLastReset time.Time

type Project struct {
	Name        string
	Description string
	URL         string
	Started     string
	Finished    string
	ImageFile   string
}

type Comment struct {
	Id        int
	Commenter string
	Email     string
	Content   string
	Timestamp time.Time
}

// Error type returned when an object is not found in database
type DatabaseError struct {
	message string
}

func (e *DatabaseError) Error() string {
	return e.message
}

// Initialises the database connection
func InitDatabase() error {
	database_url := os.Getenv("DATABASE_DIR")
	log.Info(database_url)
	var err error
	db, err = sql.Open("sqlite3", database_url)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Info("Connected to database")

	commentCount = 0
	timeSinceLastReset = time.Now()
	return nil
}

func CloseConnection() {
	if db != nil {
		db.Close()
	}
}

// fetches all projects in the database
func GetAllProjects() ([]Project, error) {
	queryString := "SELECT * FROM projects"
	rows, err := db.Query(queryString)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var proj Project
		startTime := time.Now()
		var endTime sql.NullTime

		if err := rows.Scan(&proj.Name, &proj.Description, &proj.URL, &startTime, &endTime, &proj.ImageFile); err != nil {
			log.Error(err)
			return projects, err
		}
		proj.Started = startTime.Format(time.DateOnly)
		if endTime.Valid {
			proj.Finished = endTime.Time.Format(time.DateOnly)
		} else {
			proj.Finished = "present"
		}

		projects = append(projects, proj)
	}

	if err = rows.Err(); err != nil {
		log.Error(err)
		return projects, err
	}
	return projects, nil
}

// fetches all comments from the database
func GetAllCommentIds() ([]int, error) {
	queryString := "SELECT id FROM comments ORDER BY timestamp DESC"
	rows, err := db.Query(queryString)
	if err != nil {
		return nil, &DatabaseError{err.Error()}
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Error(err)
			return nil, &DatabaseError{message: err.Error()}
		}

		ids = append(ids, id)
	}
	return ids, nil
}

// returns a single comment, obtained by its id
func GetCommentById(id int) (*Comment, error) {
	// if the comment is in cache we can return it
	if cachedComment, inCache := c.Get(strconv.Itoa(id)); inCache {
		commentPointer := cachedComment.(*Comment)
		log.Info("Loading comment from cache")
		return commentPointer, nil
	}

	queryString := "SELECT * FROM comments WHERE id = ?"
	var comment Comment
	err := db.QueryRow(queryString, id).Scan(&comment.Id, &comment.Commenter, &comment.Content, &comment.Email, &comment.Timestamp)
	if err != nil {
		log.Error(err)
		returnError := DatabaseError{message: err.Error()}
		return nil, &returnError
	}
	// add the comment to cache
	c.Set(strconv.Itoa(comment.Id), &comment, cache.DefaultExpiration)
	return &comment, nil
}

// adds a new comment to the database. To avoid api users needing to generate their own ids and risk the occasional clash,
// this is done automatically. However the input comment will not be modified, so essentially now acts as a copy.
func CreateComment(comment *Comment) (int, error) {
	// Allowing unlimited comments per day might overrun my server's capabilities.
	// I don't expect to be that popular so to prevent spam I am limiting myself to 100 comments per hour.
	if time.Since(timeSinceLastReset) > time.Hour*24 {
		commentCount = 0
		timeSinceLastReset = time.Now()
	}
	fmt.Println(commentCount)
	if commentCount >= 100 {
		return -1, &DatabaseError{message: "There have been more than 100 comments in the last day. To prevent spam this comment has been disregarded."}
	}

	queryString := "INSERT INTO comments(commenter, content, email, timestamp) VALUES (?, ?, ?, ?)"
	res, err := db.Exec(queryString, comment.Commenter, comment.Content, comment.Email, comment.Timestamp)
	if err != nil {
		log.Error(err.Error())
		return -1, &DatabaseError{message: err.Error()}
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return -1, &DatabaseError{message: "Something went wrong adding comment to database"}
	}
	if count != 1 {
		log.Error(fmt.Printf("%d comments reported as being added, not 1", count))
		return -1, &DatabaseError{message: "Something went wrong adding comment to database"}
	}

	// reassign the id of the comment to match the database id
	commentId, err := res.LastInsertId()
	if err != nil {
		return -1, &DatabaseError{message: err.Error()}
	}

	log.Info(fmt.Sprintf("Added comment with id %d to the database", commentId))
	// The input comment should be added to the cache as well.
	c.Set(strconv.Itoa(comment.Id), comment, cache.DefaultExpiration)
	commentCount++
	return int(commentId), nil
}

// edits an existing comment
func EditComment(comment *Comment) error {
	_, err := GetCommentById(comment.Id)
	if err != nil {
		return &DatabaseError{message: "The comment attempted to edit does not exist in database"}
	}
	queryString := "UPDATE comments " +
		"SET commenter=?, email=?, content=?, timestamp=? WHERE id=?"
	res, err := db.Exec(queryString, comment.Commenter, comment.Email, comment.Content, comment.Timestamp, comment.Id)
	if err != nil {
		log.Error(err.Error())
		return &DatabaseError{message: err.Error()}
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Error(err.Error())
		return &DatabaseError{message: "Something went wrong modifying the comment"}
	} else {
		log.Info(fmt.Sprintf("added entry with id %d to database", id))
	}

	// The input comment should be added to the cache as well.
	c.Set(strconv.Itoa(comment.Id), comment, cache.DefaultExpiration)
	return nil
}

// Deletes a given comment by its id
func DeleteComment(id int) error {
	queryString := "DELETE FROM comments WHERE id = ?"
	result, err := db.Exec(queryString, id)
	if err != nil {
		return &DatabaseError{err.Error()}
	}
	numAffected, err := result.RowsAffected()
	if err != nil {
		return &DatabaseError{err.Error()}
	}
	if numAffected == 0 {
		return &DatabaseError{"The comment attempted to be deleted does not exist in the database"}
	}

	c.Delete(strconv.Itoa(id))

	return nil
}
