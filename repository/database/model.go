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
	ID          uuid.UUID `repository:"id"`
	EditorID    uuid.UUID `repository:"editor_id"`
	Name        string    `repository:"name"`
	Description string    `repository:"description"`
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
	ID         uuid.UUID `repository:"id"`
	EditorID   uuid.UUID `repository:"editor_id"`
	Token      uuid.UUID `repository:"token"`
	ExpireTime time.Time `repository:"expire_time"`
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
	ID        uuid.UUID `repository:"id"`
	Password  string    `repository:"password"`
	Email     string    `repository:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
