PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE projects (
    name varchar(100) PRIMARY KEY,
    description varchar(2000),
    github_link varchar(100) NOT NULL,
    date_started date NOT NULL,
    date_finished date,
	CONSTRAINT finish_date_constraint CHECK (date_started <= date_finished) 
);
INSERT INTO projects VALUES('project 1','test project','no link','2024-10-14',NULL);
INSERT INTO projects VALUES('project 2','test project 2','no link','2024-10-15','2024-11-14');
COMMIT;

CREATE TABLE IF NOT EXISTS comments (
    id int NOT NULL PRIMARY KEY,
    commenter varchar(100) NOT NULL,
    content varchar(2000) NOT NULL,
    email varchar(100),
    timestamp TIME NOT NULL
);

INSERT INTO comments VALUES(0, "test commenter", "test content", "no email", "2024-10-20");
INSERT INTO comments VALUES(1, "test commenter two", "test content 2", "test email", "2024-12-24");