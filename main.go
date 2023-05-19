package main

import (
	"context"
	"github.com/marhycz/strv-go-newsletter/database"
	"github.com/marhycz/strv-go-newsletter/environment" */

	initFirebase "github.com/marhycz/strv-go-newsletter/database/firebase"
)

func main() {
	ctx := context.Background()

	env := &environment.Env{
		Database: database.NewConnection(ctx),
	}
	fmt.Println(env) */
	// api.Serve(env)
	initFirebase.InitSDK()
}
