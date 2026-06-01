package user_service

import (
	"context"
	"fmt"
	"net/mail"

	core_errors "github.com/wavw1/music-tracker/internal/errors"
	"github.com/wavw1/music-tracker/internal/model"
	user_domain "github.com/wavw1/music-tracker/internal/user/domain"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserServiceStruct) Register(
	ctx context.Context,
	userRequest user_domain.UserServiceInput,
) (model.User, error) {
	err := isValidEmail(userRequest.Email)
	if err != nil {
		return model.User{}, fmt.Errorf("%w: %w", core_errors.ErrInvalidEmail, err)
	}

	passwordHash, err := HashPassword(userRequest.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash: %w", err)
	}

	userModel := model.User{
		Email:        userRequest.Email,
		PasswordHash: passwordHash,
	}

	user, err := s.repo.Register(
		ctx,
		userModel,
	)
	if err != nil {
		return model.User{}, fmt.Errorf("register user repository error: %w", err)
	}

	return user, nil
}

func isValidEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
