package strv_go_newsletter

import (
	"context"
	"github.com/marhycz/strv-go-newsletter/database"
	"github.com/marhycz/strv-go-newsletter/environment"
)

func main() {
	ctx := context.Background()

	env := &environment.Env{
		Database: database.NewConnection(ctx),
	}
	// api.Serve(env)
}
