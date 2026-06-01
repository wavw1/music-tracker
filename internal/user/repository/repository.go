package user_repository

import (
	"context"

	"github.com/wavw1/music-tracker/internal/model"
)

type UserRepository interface {
	Register(
		ctx context.Context,
		userRequest model.User,
	) (model.User, error)

	GetUserByEmail(
		ctx context.Context,
		email string,
	) (model.User, error)
}
