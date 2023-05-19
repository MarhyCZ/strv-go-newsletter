package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/marhycz/strv-go-newsletter/repository/database/query"
)

func CreateEditor(ctx context.Context, querier Querier, password string, email string) (*Editor, error) {
	editor := NewEditor(password, email)
	_, err := querier.Exec(ctx, query.CreateEditor, pgx.NamedArgs{
		"id":         editor.ID,
		"password":   editor.Password,
		"email":      editor.Email,
		"created_at": editor.CreatedAt,
		"updated_at": editor.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}

	return editor, nil
}
