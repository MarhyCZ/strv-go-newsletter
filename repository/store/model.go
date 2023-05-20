package store

import "github.com/google/uuid"

type Subscription struct {
	Email        string    `json:"email"`
	Id           uuid.UUID `json:"id"`
	NewsletterID int       `json:"newsletter_id"`
}
