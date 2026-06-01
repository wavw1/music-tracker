package idea_postgresql

import (
	"context"
	"fmt"

	"github.com/wavw1/music-tracker/internal/model"
)

func (r *PostgreSql) CreateIdea(
	ctx context.Context,
	ideaRequest model.Idea,
) (model.Idea, error) {
	var ideaResponse model.Idea

	query := `
	INSERT INTO tracker.ideas (title, bpm, tonic_key, status, tags, user_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, title, bpm, tonic_key, status, tags, user_id
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		ideaRequest.Title,
		ideaRequest.Bpm,
		ideaRequest.Key,
		ideaRequest.Status,
		ideaRequest.Tags,
		ideaRequest.UserID,
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
		return model.Idea{}, fmt.Errorf("scan postgres error: %w", err)
	}

	return ideaResponse, nil
}
