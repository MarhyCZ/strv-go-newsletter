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

type renameNewsletterInput struct {
	Name string `json:"name" validate:"required"`
}

type deleteNewsletterInput struct {
	NewsletterID uuid.UUID `in:"path=newsletter_id;required;decoder=uuid"`
}

type getSubscriptionInput struct {
	NewsletterID uuid.UUID `in:"path=newsletter_id;required;decoder=uuid"`
	Email        string    `in:"path=email;required"`
}

type unsubcribeInput struct {
	SubscriptionID uuid.UUID `in:"path=subscription_id;required;decoder=uuid"`
}
