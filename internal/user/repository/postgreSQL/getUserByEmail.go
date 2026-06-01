package user_postgresql

import (
	"context"
	"fmt"

	"github.com/wavw1/music-tracker/internal/model"
)

func (r *PostgreSql) GetUserByEmail(
	ctx context.Context,
	email string,
) (model.User, error) {
	var user model.User

	query := `
	SELECT id, email, password_hash, created_at FROM tracker.users
	WHERE email = $1
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		email,
	)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		return model.User{}, fmt.Errorf("query user error: %w", err)
	}

	return user, nil
}
