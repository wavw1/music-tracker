package user_handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wavw1/music-tracker/internal/dto"
	core_errors "github.com/wavw1/music-tracker/internal/errors"
	user_domain "github.com/wavw1/music-tracker/internal/user/domain"
)

func (h *UserHandler) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var loginModel = user_domain.UserServiceInput{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}

	user, err := h.service.Login(
		c.Request.Context(),
		loginModel,
	)
	if err != nil {
		if errors.Is(err, core_errors.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.LoginResponse{
		Token: user.Token,
		ID:    user.UserID,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, response)
}
