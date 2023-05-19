package environment

import (
	"github.com/marhycz/strv-go-newsletter/repository/database"
)

type Env struct {
	Database *database.Database
}
