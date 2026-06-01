package idea_postgresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	core_errors "github.com/wavw1/music-tracker/internal/errors"
	"github.com/wavw1/music-tracker/internal/model"
)

func (r *PostgreSql) GetIdea(
	ctx context.Context,
	id int,
) (model.Idea, error) {
	var ideaResponse model.Idea

	query := `
	SELECT id, title, bpm, tonic_key, status, tags, user_id FROM tracker.ideas
	WHERE id=$1
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		id,
	)
	err := row.Scan(
		&ideaResponse.ID,
		&ideaResponse.Title,
		&ideaResponse.Bpm,
		&ideaResponse.Key,
		&ideaResponse.Status,
		&ideaResponse.Tags,
		&ideaResponse.UserID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Idea{}, core_errors.ErrNotFound
		}

		return model.Idea{}, fmt.Errorf("scan postgres error: %w", err)
	}

	return ideaResponse, nil
}
