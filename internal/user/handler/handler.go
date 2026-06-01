package user_handler

import user_service "github.com/wavw1/music-tracker/internal/user/service"

type UserHandler struct {
	service user_service.UserService
}

func NewUserHandler(s user_service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}
