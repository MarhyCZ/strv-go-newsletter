package database

import (
	"time"

	"github.com/google/uuid"
)

func NewNewsletter(editorID uuid.UUID, name string, description string) *Newsletter {
	now := time.Now()
	return &Newsletter{
		ID:          uuid.New(),
		EditorID:    editorID,
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

type Newsletter struct {
	ID          uuid.UUID `db:"id"`
	EditorID    uuid.UUID `db:"editor_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func NewPasswordReset(editorID uuid.UUID, expireTime time.Time) *PasswordReset {
	return &PasswordReset{
		ID:         uuid.New(),
		EditorID:   editorID,
		Token:      uuid.New(),
		ExpireTime: expireTime,
	}
}

type PasswordReset struct {
	ID         uuid.UUID `db:"id"`
	EditorID   uuid.UUID `db:"editor_id"`
	Token      uuid.UUID `db:"token"`
	ExpireTime time.Time `db:"expire_time"`
}

func NewEditor(password string, email string) *Editor {
	now := time.Now()
	return &Editor{
		ID:        uuid.New(),
		Password:  password,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type Editor struct {
	ID        uuid.UUID `db:"id"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
