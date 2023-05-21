package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/marhycz/strv-go-newsletter/repository/database"

	"github.com/go-chi/chi/v5"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserDoesntExists  = errors.New("user does not exist")
)

func (rest *Rest) routeEditor(r *chi.Mux) {
	// RESTy routes for "articles" resource
	r.Route("/login", func(r chi.Router) {
		r.Post("/", rest.login) // POST /editor
	})
	r.Route("/signup", func(r chi.Router) {
		r.Post("/", rest.signup) // POST /editor
	})
	r.Route("/resetpassword", func(r chi.Router) {
		r.Post("/", rest.resetPassword) // POST /editor
	})
	r.Route("/getEditors", func(r chi.Router) {
		r.Get("/", rest.getEditors) // POST /editor
	})
}

func (rest *Rest) signup(w http.ResponseWriter, r *http.Request) {

	newEditorInput := database.NewEditorInput{}

	if err := parseRequestBody(r, &newEditorInput); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	db := rest.env.Database.Database

	if exists, _ := database.GetEditorByEmail(ctx, db, newEditorInput.Email); exists != nil {
		w.WriteHeader(http.StatusConflict)
		panic(ErrUserAlreadyExists)
	}

	editor, err := database.CreateEditor(ctx, db, newEditorInput.Password, newEditorInput.Email)
	if err != nil {
		panic(err)
	}

	response := fmt.Sprintf("The editor %s was successfully created.", editor.Email)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (rest *Rest) getEditors(w http.ResponseWriter, r *http.Request) {

	db := rest.env.Database.Database

	editors, err := database.ListEditors(r.Context(), db)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	for i := 0; i <= len(editors); i++ {
		w.Write([]byte(editors[i].Email + "\n"))
	}
}

func (rest *Rest) login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("here is jwt token"))
}

func (rest *Rest) resetPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("reset email sent succesfully"))
}
