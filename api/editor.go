package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
		w.Write([]byte("wrong login credentials"))
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: LoginInput.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.Write([]byte(tokenString))

}

func (rest *Rest) resetPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("reset email sent succesfully"))
}
