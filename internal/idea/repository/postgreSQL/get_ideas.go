package idea_postgresql

import (
	"context"
	"fmt"

	"github.com/wavw1/music-tracker/internal/model"
)

func (r *PostgreSql) GetIdeas(
	ctx context.Context,
	limit int,
	offset int,
	userID int64,
) ([]model.Idea, error) {
	var ideas []model.Idea

	query, args := queryConstructor(limit, offset, userID)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query postgres error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var idea model.Idea

		err := rows.Scan(
			&idea.ID,
			&idea.Title,
			&idea.Bpm,
			&idea.Key,
			&idea.Status,
			&idea.Tags,
			&idea.UserID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan postgres error: %w", err)
		}

		ideas = append(ideas, idea)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ideas, nil
}

func queryConstructor(
	limit int,
	offset int,
	userID int64,
) (string, []any) {
	var query string
	var args []any

	if limit != 0 && offset != 0 {
		query = `
	SELECT id, title, bpm, tonic_key, status, tags, user_id FROM tracker.ideas
	WHERE user_id = $1
	ORDER BY id ASC
	LIMIT $2 OFFSET $3
	`

		args = append(args, userID, limit, offset)

	} else if limit != 0 {
		query = `
	SELECT id, title, bpm, tonic_key, status, tags, user_id FROM tracker.ideas
	WHERE user_id = $1
	ORDER BY id ASC
	LIMIT $2
	`

		args = append(args, userID, limit)

	} else if offset != 0 {
		query = `
	SELECT id, title, bpm, tonic_key, status, tags, user_id FROM tracker.ideas
	WHERE user_id = $1
	ORDER BY id ASC
	OFFSET $2
	`

		args = append(args, userID, offset)

	} else {
		query = `
	SELECT id, title, bpm, tonic_key, status, tags, user_id FROM tracker.ideas
	WHERE user_id = $1
	ORDER BY id ASC`

		args = append(args, userID)
	}

	return query, args
}
