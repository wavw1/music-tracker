package user_postgresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	core_errors "github.com/wavw1/music-tracker/internal/errors"
	"github.com/wavw1/music-tracker/internal/model"
)

func (r *PostgreSql) Register(
	ctx context.Context,
	userRequest model.User,
) (model.User, error) {
	var userResponse model.User

	query := `
	INSERT INTO tracker.users (email, password_hash)
	VALUES ($1, $2)
	RETURNING id, email, created_at
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		userRequest.Email,
		userRequest.PasswordHash,
	)
	err := row.Scan(
		&userResponse.ID,
		&userResponse.Email,
		&userResponse.CreatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return model.User{}, fmt.Errorf("%w: %w", core_errors.ErrEmailExists, err)
		}

		return model.User{}, fmt.Errorf("scan postgres error: %w", err)
	}

	return userResponse, nil
}
