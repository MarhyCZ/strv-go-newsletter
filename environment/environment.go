package environment

import (
	"github.com/marhycz/strv-go-newsletter/repository/database"
	"github.com/marhycz/strv-go-newsletter/repository/store"
)

type Env struct {
	Database *database.Database
	Store    *store.Store
}
