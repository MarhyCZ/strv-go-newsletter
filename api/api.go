package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/marhycz/strv-go-newsletter/api/routes"
	"github.com/marhycz/strv-go-newsletter/environment"
	"net/http"
	"os"
	"time"
)

func Serve(env *environment.Env) {
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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	routes.RouteEditor(r)
	http.ListenAndServe(":"+os.Getenv("API_PORT"), r)
}
