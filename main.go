package main

import (
	/* "fmt"

	"github.com/marhycz/strv-go-newsletter/database"
	"github.com/marhycz/strv-go-newsletter/environment" */

	initFirebase "github.com/marhycz/strv-go-newsletter/database/firebase"
)

func main() {
	/* env := &environment.Env{
		Database: database.NewConnection(),
	}
	fmt.Println(env) */
	// api.Serve(env)
	initFirebase.InitSDK()
}
