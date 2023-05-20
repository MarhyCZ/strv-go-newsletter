package database

import (
	"context"
	"embed"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

//go:embed migrations/*.sql
var migrationFs embed.FS

type Database struct {
	database *pgxpool.Pool
}

// https://github.com/jackc/pgx/issues/1188
type Querier interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...any) (pgx.Rows, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}

func NewConnection(ctx context.Context) *Database {
	var (
		host     = os.Getenv("POSTGRES_HOST")
		port     = os.Getenv("POSTGRES_PORT")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, password, host, port, dbname)

	db, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to repository: %v\n", err)
		os.Exit(1)
	}
	// defer db.Close()

	err = db.Ping(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to DB!")

	d := &Database{
		database: db,
	}

	return d
}
