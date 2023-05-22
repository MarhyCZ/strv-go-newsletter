package store

type Subscription struct {
	Email        string    `json:"email"`
	Id           string    `json:"id"`
	Newsletter_id int       `json:"newsletter_id"`
}
