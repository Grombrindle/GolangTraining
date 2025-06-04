package Sql

import (
	sql "database/sql"
	"log"
	"time"

	_ "github.com/glebarez/go-sqlite"
	// _ "go-sql-driver/mysql"
	// "errors"
)

func DatabaseTest() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// db, err := sql.Open("mysql", "root:1234@(127.0.0.1:3306)/dbname?parseTime=true")
	db, err := sql.Open("sqlite", "./database.db")

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	query := `
		CREATE TABLE IF NOT EXISTS USERS (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME
		);`
	// defer db.Close()
	_, err2 := db.Exec(query)
	if err2 != nil {
		panic(err2)
	}

	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO USERS (username,password,created_at) VALUES (?,?,?)`, username, password, createdAt)

	if err != nil {
		panic(err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	print(userID)

	var (
		id         int
		username2  string
		password2  string
		createdAt2 time.Time
	)
	query2 := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err = db.QueryRow(query2, 4).Scan(&id, &username2, &password2, &createdAt2)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
		panic(err)
	}

	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}
	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			panic(err)
		}
		users = append(users, u)
	}
	err3 := rows.Err()
	if err3 != nil {
		panic(err3)
	}
	print(users)

}
