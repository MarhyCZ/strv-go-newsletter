package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/marhycz/strv-go-newsletter/repository/database"

	"github.com/go-chi/chi/v5"
)

var (
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrUserDoesntExists     = errors.New("user does not exist")
	ErrUserWrongCredentials = errors.New("wrong login credentials")
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
	r.Route("/logout", func(r chi.Router) {
		r.Get("/", rest.logout) // POST /editor
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

	passwordHash, err := HashPassword(newEditorInput.Password)
	if err != nil {
		panic(err)
	}

	editor, err := database.CreateEditor(ctx, db, passwordHash, newEditorInput.Email)
	if err != nil {
		panic(err)
	}

	response := fmt.Sprintf("The editor %s was successfully created.", editor.Email)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (rest *Rest) getEditors(w http.ResponseWriter, r *http.Request) {
	if c, err := AuthToken(r); err != nil {
		w.WriteHeader(c)
		w.Write([]byte(strconv.Itoa(c) + ": " + http.StatusText(c)))
		return
	}

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

	LoginInput := database.LoginInput{}
	if err := parseRequestBody(r, &LoginInput); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db := rest.env.Database.Database

	e, err := database.GetEditorByEmail(r.Context(), db, LoginInput.Email)
	if err != nil {
		panic(err)
	}
	if !CheckPasswordHash(LoginInput.Password, e.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(strconv.Itoa(http.StatusUnauthorized) + ": " + http.StatusText(http.StatusUnauthorized)))
		w.Write([]byte("\n" + "Wrong login credentials"))
		return
	}

	token, expDate, err := CreateNewJWT(LoginInput.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(strconv.Itoa(http.StatusInternalServerError) + ": " + http.StatusText(http.StatusInternalServerError)))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expDate,
	})

	w.Write([]byte("JTW token: " + token))

}

func (rest *Rest) resetPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("reset email sent succesfully"))
}

func (rest *Rest) logout(w http.ResponseWriter, r *http.Request) {
	// immediately clear the token cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("logout successful"))
}
