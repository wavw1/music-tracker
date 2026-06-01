package user_service

import (
	"context"

	"github.com/wavw1/music-tracker/internal/model"
	user_domain "github.com/wavw1/music-tracker/internal/user/domain"
	user_repository "github.com/wavw1/music-tracker/internal/user/repository"
)

type UserServiceStruct struct {
	repo user_repository.UserRepository
}

func NewUserServiceStruct(
	repo user_repository.UserRepository,
) *UserServiceStruct {
	return &UserServiceStruct{
		repo: repo,
	}
}

type UserService interface {
	Register(
		ctx context.Context,
		userRequest user_domain.UserServiceInput,
	) (model.User, error)

	Login(
		ctx context.Context,
		userRequest user_domain.UserServiceInput,
	) (user_domain.LoginResult, error)
}
