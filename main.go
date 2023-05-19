package main

import (
	"context"
	"fmt"
	"github.com/marhycz/strv-go-newsletter/environment"
	"github.com/marhycz/strv-go-newsletter/repository/database"

	initFirebase "github.com/marhycz/strv-go-newsletter/repository/firebase"
)

func main() {
	ctx := context.Background()

	env := &environment.Env{
		Database: database.NewConnection(ctx),
	}
	fmt.Println(env)
	// api.Serve(env)
	initFirebase.InitSDK()
}
