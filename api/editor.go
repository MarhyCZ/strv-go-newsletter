package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (rest *Rest) routeEditor(r *chi.Mux) {
	// RESTy routes for "articles" resource
	r.Route("/login", func(r chi.Router) {
		r.Post("/", rest.login) // POST /editor
	})
	r.Route("/signup", func(r chi.Router) {
		r.Post("/signup", rest.signup) // POST /editor
	})
	r.Route("/resetpassword", func(r chi.Router) {
		r.Post("/", rest.resetPassword) // POST /editor
	})
}

func (rest *Rest) signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func (rest *Rest) login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("here is jwt token"))
}

func (rest *Rest) resetPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("reset email sent succesfully"))
}
