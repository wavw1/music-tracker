package idea_postgresql

import (
	"context"
	"fmt"

	core_errors "github.com/wavw1/music-tracker/internal/errors"
)

func (r *PostgreSql) DeleteIdea(
	ctx context.Context,
	ideaID int,
) error {
	query := `
	DELETE FROM tracker.ideas
	WHERE id=$1
	`

	tag, err := r.pool.Exec(
		ctx,
		query,
		ideaID,
	)
	if err != nil {
		return fmt.Errorf("exec postgres error: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return core_errors.ErrNotFound
	}

	return nil
}
