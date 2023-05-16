package database

import (
	"database/sql"
	"fmt"
)

type Database struct {
	database *sql.DB
}

func NewConnection() *Database {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "your-password"
		dbname   = "strvdb"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	d := &Database{
		database: db,
	}

	return d
}
