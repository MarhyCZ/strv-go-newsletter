package strv_go_newsletter

import (
	"github.com/marhycz/strv-go-newsletter/database"
	"github.com/marhycz/strv-go-newsletter/environment"
)

func main() {
	env := &environment.Env{
		Database: database.NewConnection(),
	}
	// api.Serve(env)
}
