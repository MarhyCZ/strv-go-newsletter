package database

import (
	"context"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/marhycz/strv-go-newsletter/repository/database/query"
)

func ResetPasswordRequest(ctx context.Context, querier Querier, editorID uuid.UUID, expireTime time.Time) (*PasswordReset, error) {
	passwordReset := NewPasswordReset(editorID, expireTime)
	_, err := querier.Exec(ctx, query.CreatePasswordReset, pgx.NamedArgs{
		"id":          passwordReset.ID,
		"editor_id":   passwordReset.EditorID,
		"token":       passwordReset.Token,
		"expire_time": passwordReset.ExpireTime,
	})
	if err != nil {
		return nil, err
	}
	return passwordReset, nil
}

func UpdateEditorPassword(ctx context.Context, querier Querier, editorID uuid.UUID, newPassword string) error {
	_, err := querier.Exec(ctx, query.ResetPassword, pgx.NamedArgs{
		"id":         editorID,
		"password":   newPassword,
		"updated_at": time.Now(),
	})
	return err
}

func GetResetPasswordRequest(ctx context.Context, querier Querier, token string) (*PasswordReset, error) {
	var passwordReset PasswordReset
	err := pgxscan.Get(ctx, querier, &passwordReset, query.GetPasswordReset, pgx.NamedArgs{
		"token": token,
	})

	if err != nil {
		return nil, err
	}
	return &passwordReset, nil
}
