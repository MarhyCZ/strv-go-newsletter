package api

import (
	"github.com/marhycz/strv-go-newsletter/environment"
	"net/http"
)

func Serve(env environment.Env) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
