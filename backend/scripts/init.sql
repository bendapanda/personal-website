CREATE TABLE IF NOT EXISTS projects (
    name varchar(100) PRIMARY KEY,
    description varchar(2000),
    github_link varchar(100) NOT NULL,
    date_started date NOT NULL,
    date_finished date,
    image_file varchar(200),
	CONSTRAINT finish_date_constraint CHECK (date_started <= date_finished) 
);

CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY,
    commenter varchar(100) NOT NULL,
    content varchar(2000) NOT NULL,
    email varchar(100),
    timestamp date NOT NULL
);

INSERT INTO projects(name, description, github_link, date_started, date_finished, image_file) VALUES 
    ('Handwriting Recognition System',
    'This is a handwriting recognition system, build from scratch, by me in python. 
    In theory it takes images of handwriting and outputs the text it contains, although it is very much 
    still a work in progress',
    'https://github.com/bendapanda/note-monkey',
    '2024-03-16',
    NULL,
    "resources/note_money.png");

INSERT INTO projects(name, description, github_link, date_started, date_finished, image_file) VALUES 
    ('Personal Website',
    'This is the repository for the website you are on right now!',
    'https://github.com/bendapanda/personal-website',
    '2024-08-28',
    NULL,
    "resources/personal_website.png");

INSERT INTO comments(commenter, content, email, timestamp) VALUES
    (
    "Ben Shirley",
    "Wow Ben! This is an awesome website!",
    "benshirley04@gmail.com",
    "2024-10-20");
