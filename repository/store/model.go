package store

import "github.com/google/uuid"

type Subscription struct {
	Email         string    `json:"email"`
	Id            string    `json:"id"`
	Newsletter_id uuid.UUID `json:"newsletter_id"`
}
