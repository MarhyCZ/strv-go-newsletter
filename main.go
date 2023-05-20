package main

import (
	"context"
	"github.com/marhycz/strv-go-newsletter/api"
	"github.com/marhycz/strv-go-newsletter/environment"
	"github.com/marhycz/strv-go-newsletter/repository/database"
	"github.com/marhycz/strv-go-newsletter/repository/store"
	"log"
	"net/http"
	"os"
)

func main() {
	// Context should  be passed directly to functions, not stored in structs
	// https://pkg.go.dev/context#Background
	ctx := context.Background()

	env := &environment.Env{
		Database: database.NewConnection(ctx),
		Store:    store.NewConnection(ctx),
	}

	controller := api.NewController(env)
	if err := http.ListenAndServe(":"+os.Getenv("API_PORT"), controller); err != nil {
		log.Fatal(err.Error())
	}
}
