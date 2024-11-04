PRAGMA foreign_keys=OFF;

DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS comments;

CREATE TABLE IF NOT EXISTS projects (
    name varchar(100) PRIMARY KEY,
    description varchar(2000),
    github_link varchar(100) NOT NULL,
    date_started date NOT NULL,
    date_finished date,
    image_file varchar(200),
	CONSTRAINT finish_date_constraint CHECK (date_started <= date_finished) 
);
INSERT INTO projects VALUES('project 1','test project','no link','2024-10-14',NULL, "test_image.png");
INSERT INTO projects VALUES('project 2','test project 2','no link','2024-10-15','2024-11-14', "test_image.png");

CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY, 
    commenter varchar(100) NOT NULL,
    content varchar(2000) NOT NULL,
    email varchar(100),
    timestamp date NOT NULL
);

INSERT INTO comments(commenter, content, email, timestamp) VALUES("test commenter", "test content", "no email", "2024-10-20");
INSERT INTO comments(commenter, content, email, timestamp) VALUES("test commenter two", "test content 2", "test email", "2024-12-24");