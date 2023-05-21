package environment

import (
	"github.com/marhycz/strv-go-newsletter/repository/database"
	"github.com/marhycz/strv-go-newsletter/repository/store"
	"github.com/marhycz/strv-go-newsletter/service"
)

type Env struct {
	Database *database.Database
	Store    *store.Store
	Service  *service.Service
}
