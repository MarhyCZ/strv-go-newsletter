package database

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
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

func ListEditors(ctx context.Context, querier Querier) ([]Editor, error) {
	var editors []Editor
	if err := pgxscan.Select(ctx, querier, &editors, query.ListEditors); err != nil {
		return nil, err
	}
	return editors, nil
}

func GetEditorByEmail(ctx context.Context, querier Querier, email string) (*Editor, error) {
	var editor Editor
	err := pgxscan.Get(ctx, querier, &editor, query.GetEditor, pgx.NamedArgs{
		"email": email,
	})

	if err != nil {
		return nil, err
	}
	return &editor, nil
}
