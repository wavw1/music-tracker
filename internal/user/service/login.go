package user_service

import (
	"context"
	"fmt"

	"github.com/wavw1/music-tracker/internal/auth"
	core_errors "github.com/wavw1/music-tracker/internal/errors"
	"github.com/wavw1/music-tracker/internal/model"
	user_domain "github.com/wavw1/music-tracker/internal/user/domain"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserServiceStruct) Login(
	ctx context.Context,
	userRequest user_domain.UserServiceInput,
) (user_domain.LoginResult, error) {
	requestUserModel := model.User{
		Email: userRequest.Email,
	}

	repoUser, err := s.repo.GetUserByEmail(
		ctx,
		requestUserModel.Email,
	)
	if err != nil {
		return user_domain.LoginResult{}, core_errors.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(repoUser.PasswordHash),
		[]byte(userRequest.Password),
	)
	if err != nil {
		return user_domain.LoginResult{}, core_errors.ErrInvalidCredentials
	}

	token, err := auth.GenerateToken(repoUser)
	if err != nil {
		return user_domain.LoginResult{}, fmt.Errorf("failed to generate token: %w", err)
	}

	loginResult := user_domain.LoginResult{
		Token:  token,
		UserID: repoUser.ID,
		Email:  repoUser.Email,
	}

	return loginResult, nil
}
