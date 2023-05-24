package database

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/marhycz/strv-go-newsletter/repository/database/query"
	"time"
)

func CreateNewsletter(ctx context.Context, querier Querier, editorID uuid.UUID, name string, description string) (*Newsletter, error) {
	newsletter := NewNewsletter(editorID, name, description)
	_, err := querier.Exec(ctx, query.CreateNewsletter, pgx.NamedArgs{
		"id":          newsletter.ID,
		"editor_id":   newsletter.EditorID,
		"name":        newsletter.Name,
		"description": newsletter.Description,
		"created_at":  newsletter.CreatedAt,
		"updated_at":  newsletter.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}
	return newsletter, nil
}

func ListEditorNewsletters(ctx context.Context, querier Querier, editorID uuid.UUID) ([]Newsletter, error) {
	var newsletters []Newsletter
	if err := pgxscan.Select(ctx, querier, &newsletters, query.ListNewsletters, pgx.NamedArgs{
		"editor_id": editorID,
	}); err != nil {
		return nil, err
	}
	return newsletters, nil
}

func RenameNewsletter(ctx context.Context, querier Querier, newsletterID uuid.UUID, name string) error {
	_, err := querier.Exec(ctx, query.RenameNewsletter, pgx.NamedArgs{
		"id":         newsletterID,
		"name":       name,
		"updated_at": time.Now(),
	})
	return err
}

func DeleteNewsletter(ctx context.Context, querier Querier, newsletterID uuid.UUID) error {
	_, err := querier.Exec(ctx, query.DeleteNewsletter, pgx.NamedArgs{
		"id": newsletterID,
	})
	return err
}
