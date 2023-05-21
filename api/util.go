package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
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
