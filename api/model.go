package api

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type newEditorInput struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type loginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ResetPasswordInput struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type createNewsletterInput struct {
	EditorID    uuid.UUID `json:"editor_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
}

type deleteNewsletterInput struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type listEditorNewslettersInput struct {
	EditorID uuid.UUID `json:"editor_id" validate:"required"`
}
