package api

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type claims struct {
	Username string    `json:"username"`
	EditorID uuid.UUID `json:"editor_id"`
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

type resetPasswordInput struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type createNewsletterInput struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type deleteNewsletterInput struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type renameNewsletterInput struct {
	ID   uuid.UUID `json:"id" validate:"required"`
	Name string    `json:"description" validate:"required"`
}

type listEditorNewslettersInput struct {
	EditorID uuid.UUID `json:"editor_id" validate:"required"`
}
