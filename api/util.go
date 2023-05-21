package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
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

func AuthToken(r *http.Request) (int, error) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return http.StatusUnauthorized, err
		}
		return http.StatusBadRequest, err
	}

	tknStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return http.StatusUnauthorized, err
		}
		return http.StatusBadRequest, err
	}
	if !tkn.Valid {
		return http.StatusBadRequest, err
	}
	return http.StatusAccepted, nil
}

func CreateNewJWT(email string) (string, time.Time, error) {

	//30 minutes
	expirationTime := time.Now().Add(30 * time.Minute)

	claims := &Claims{
		Username: email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", expirationTime, err
	}

	return tokenString, expirationTime, nil
}
