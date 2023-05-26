package api

import (
	"encoding/json"
	"github.com/marhycz/strv-go-newsletter/repository/database"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	validate = validator.New()
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

func authToken(r *http.Request) (*claims, int) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, http.StatusUnauthorized
		}
		return nil, http.StatusBadRequest
	}

	tknStr := c.Value
	claims := &claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, http.StatusUnauthorized
		}
		return nil, http.StatusBadRequest
	}
	if !tkn.Valid {
		return claims, http.StatusBadRequest
	}
	return claims, http.StatusAccepted
}

func CreateNewJWT(editor *database.Editor) (string, time.Time, error) {

	//30 minutes
	expirationTime := time.Now().Add(30 * time.Minute)

	claims := &claims{
		Username: editor.Email,
		EditorID: editor.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", expirationTime, err
	}

	return tokenString, expirationTime, nil
}

func Validate(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	validEmail := regexp.MustCompile(emailRegex).MatchString(email)
	return validEmail
}
