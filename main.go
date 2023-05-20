package main

import (
	"context"
	"fmt"
	"github.com/marhycz/strv-go-newsletter/api"
	"github.com/marhycz/strv-go-newsletter/environment"
	"github.com/marhycz/strv-go-newsletter/repository/database"

	"github.com/marhycz/strv-go-newsletter/repository/store"
)

func main() {
	// Context should  be passed directly to functions, not stored in structs
	// https://pkg.go.dev/context#Background
	ctx := context.Background()

	env := &environment.Env{
		Database: database.NewConnection(ctx),
		Store:    store.NewConnection(ctx),
	}
	fmt.Println(env)
	api.Serve(env)
	env.Store.GetSubscriptions(ctx)
}
