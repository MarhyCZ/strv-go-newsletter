package database

import (
	"database/sql"
	"fmt"
	"os"
)

type Database struct {
	database *sql.DB
}

func NewConnection() *Database {
	var (
		host     = os.Getenv("POSTGRE_HOST")
		port     = os.Getenv("POSTGRE_PORT")
		user     = os.Getenv("POSTGRE_USER")
		password = os.Getenv("POSTGRE_PASS")
		dbname   = os.Getenv("POSTGRE_DB")
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
