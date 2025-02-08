package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./grit-space.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	createPostsTable := `CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		category TEXT NOT NULL
	);`

	createCommentsTable := `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		FOREIGN KEY (post_id) REFERENCES posts(id)
	);`

	createSessionsTable := `CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		session_token TEXT NOT NULL,
		expires_at DATETIME NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	createLikesTable := `CREATE TABLE IF NOT EXISTS likes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (post_id) REFERENCES posts(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	createDislikesTable := `CREATE TABLE IF NOT EXISTS dislikes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (post_id) REFERENCES posts(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(createPostsTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(createCommentsTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(createSessionsTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(createLikesTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(createDislikesTable)
	if err != nil {
		log.Fatal(err)
	}
}
