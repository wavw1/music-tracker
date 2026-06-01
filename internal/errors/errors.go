package core_errors

import (
	"errors"
)

var (
	ErrInvalidBPM         = errors.New("Bpm should be greater than 50 and less than 550")
	ErrNotFound           = errors.New("Idea not exist")
	ErrNoFieldsToUpdate   = errors.New("No fields to update")
	ErrEmailExists        = errors.New("Email already exists")
	ErrInvalidEmail       = errors.New("Invalid email")
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrEmptySecret        = errors.New("empty JWT secret key")
	ErrInvalidToken       = errors.New("Invalid token")
	ErrTokenRequired      = errors.New("API token required")
	ErrTokenFormat        = errors.New("invalid token format")
	ErrForbidden          = errors.New("InsufficientPermissions")
)
