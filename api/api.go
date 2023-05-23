package api

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/marhycz/strv-go-newsletter/environment"
)

type Rest struct {
	*chi.Mux
	env *environment.Env
}

func NewController(env *environment.Env) *Rest {
	c := &Rest{
		env: env,
	}
	c.initRouter()
	return c
}

func (rest *Rest) initRouter() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	rest.routeEditor(r)
	rest.routeSubscriptions(r)
	rest.routeIssues(r)
	rest.Mux = r
}
