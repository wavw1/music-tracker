package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	core_errors "github.com/wavw1/music-tracker/internal/errors"
	"github.com/wavw1/music-tracker/internal/model"
)

func GenerateToken(user model.User) (string, error) {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		return "", core_errors.ErrEmptySecret
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(key))
}
