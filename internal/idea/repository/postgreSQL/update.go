package idea_postgresql

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	core_errors "github.com/wavw1/music-tracker/internal/errors"
	idea_domain "github.com/wavw1/music-tracker/internal/idea/domain"
	"github.com/wavw1/music-tracker/internal/model"
)

func (r *PostgreSql) UpdateIdea(
	ctx context.Context,
	ideaID int,
	patch idea_domain.IdeaPatch,
) (
	model.Idea,
	error,
) {
	setParts := []string{}
	args := []any{}
	argID := 1

	if patch.Title != nil {
		setParts = append(setParts, fmt.Sprintf("title = $%d", argID))
		args = append(args, *patch.Title)
		argID++
	}

	if patch.Bpm != nil {
		setParts = append(setParts, fmt.Sprintf("bpm = $%d", argID))
		args = append(args, *patch.Bpm)
		argID++
	}

	if patch.Key != nil {
		setParts = append(setParts, fmt.Sprintf("tonic_key = $%d", argID))
		args = append(args, *patch.Key)
		argID++
	}

	args = append(args, ideaID)

	query := fmt.Sprintf(`
	UPDATE tracker.ideas
	SET %s
	WHERE id=$%d
	RETURNING id, title, bpm, tonic_key, status, tags, user_id
	`, strings.Join(setParts, ","),
		argID,
	)

	if len(setParts) == 0 {
		return model.Idea{}, core_errors.ErrNoFieldsToUpdate
	}

	var ideaResponse model.Idea

	row := r.pool.QueryRow(
		ctx,
		query,
		args...,
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
