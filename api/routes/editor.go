package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/marhycz/strv-go-newsletter/repository/database"
	"net/http"
)

func RouteEditor(r *chi.Mux) {
	// RESTy routes for "articles" resource
	r.Route("/login", func(r chi.Router) {
		r.Post("/", login) // POST /editor
	})
	r.Route("/signup", func(r chi.Router) {
		r.Post("/signup", signup) // POST /editor
	})
	r.Route("/resetpassword", func(r chi.Router) {
		r.Post("/", resetPassword) // POST /editor
	})
}

func signup(w http.ResponseWriter, r *http.Request) {
	database.CreateEditor()
	w.Write([]byte("hi"))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("here is jwt token"))
}

func resetPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("reset email sent succesfully"))
}
