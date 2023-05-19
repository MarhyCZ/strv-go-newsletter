package database

import (
	"github.com/google/uuid"
	"time"
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
	ID          uuid.UUID `database:"id"`
	EditorID    uuid.UUID `database:"editor_id"`
	Name        string    `database:"name"`
	Description string    `database:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func NewPasswordReset(editorID uuid.UUID, expireTime time.Time) *PasswordReset {
	return &PasswordReset{
		ID:         uuid.New(),
		EditorID:   editorID,
		ExpireTime: expireTime,
	}
}

type PasswordReset struct {
	ID         uuid.UUID `database:"id"`
	EditorID   uuid.UUID `database:"editor_id"`
	ExpireTime time.Time `database:"expire_time"`
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
	ID        uuid.UUID `database:"id"`
	Password  string    `database:"password"`
	Email     string    `database:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
