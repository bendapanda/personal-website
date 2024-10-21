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
	Email     sql.NullString
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
	database_url := os.Getenv("DATABASE_URL")
	log.Info(database_url)
	var err error
	db, err = sql.Open("sqlite3", database_url)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Info("Connected to database")
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
	queryString := "SELECT id FROM comments ORDER BY timestamp"
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

// adds a new comment to the database
func CreateComment(comment *Comment) error {
	_, err := GetCommentById(comment.Id)
	if err == nil {
		return &DatabaseError{message: fmt.Sprintf("A comment with id %d already exists in the database!", comment.Id)}
	}

	queryString := "INSERT INTO comments(id, commenter, content, email, timestamp) VALUES (?, ?, ?, ?, ?)"
	res, err := db.Exec(queryString, comment.Id, comment.Commenter, comment.Content, comment.Email, comment.Timestamp)
	if err != nil {
		log.Error(err.Error())
		return &DatabaseError{message: err.Error()}
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Error(err.Error())
		return &DatabaseError{message: "Something went wrong adding comment to database"}
	}
	log.Info(fmt.Sprintf("added entry with id %d to database", id))

	// The input comment should be added to the cache as well.
	c.Set(strconv.Itoa(comment.Id), comment, cache.DefaultExpiration)
	return nil
}

// edits an existing comment
func EditComment(Comment) error {
	return nil
}

// Deletes a given comment by its id
func DeleteComment(int) error {
	return nil
}
