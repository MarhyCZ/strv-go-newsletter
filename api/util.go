package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var (
	validate = validator.New()
	pepper   []byte
)

const (
	bcryptMaxPasswordLength = 72
)

func parseRequestBody(r *http.Request, target any) error {
	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		return err
	}
	if err := validate.Struct(target); err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
