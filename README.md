# Personal website

This is my full-stack personal website, built with react and go, using sqlite as a backend.
It's definitely a bit overkill to use all this for a website that realistically could be static,
but I figure this can function as a bit of a portfolio-piece too.

This is still a work in progress, so it's not live anywhere, but hopefully it will be soon!

## About the project

When making this website I wanted to show off my skills. I've kept the tech stack pretty fundamental, in order to
focus on the basics. React is primarily used on the frontend to dynamically update content in response to api requests.
The backend uses a restful api to provide services to the frontend. This allows it to fetch my projects from a database, 
and then render them on the website. This also allows me to have a fully functional comments section!

## Requirements

The required installations for this project are Go, npm, sqlite3.
To run the website, do the following:

1. First, a database needs to be created. This is not included in the repo, but a template sql file, init.sql,
    is located in "/backend/scripts". Once the database has been created, set the environment variable
    `DATABASE_DIRECTORY=/path/to/database.db`.
2. Now, spin up the backend with `go run server.go` (cd into backend). There should be no error in the logs.
    Logs can also be sent to a file, by setting the environment variable `LOG_FILE` accordingly.
3. Finally, the frontend can be started, by first running `npm i` to install dependencies, and then running `npm start`.

Following these steps a development server should be started!

## Future TODOs

There are several things I would like to add to this project:
I think it would be cool to have a login system for the comments section, as this would allow users to edit and delete
their own comments.

I would like to add a "fun-zone" to show off the other cool projects that I have!

Probably most importantly, this site is not designed with mobile users in mind at all! I definitely need to add a mobile view.