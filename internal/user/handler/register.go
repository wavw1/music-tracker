package user_handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wavw1/music-tracker/internal/dto"
	core_errors "github.com/wavw1/music-tracker/internal/errors"
	user_domain "github.com/wavw1/music-tracker/internal/user/domain"
)

func (h *UserHandler) Register(c *gin.Context) {
	var request dto.UserDTORequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var registerModel = user_domain.UserServiceInput{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.service.Register(
		c.Request.Context(),
		registerModel,
	)
	if err != nil {
		if errors.Is(err, core_errors.ErrInvalidEmail) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if errors.Is(err, core_errors.ErrEmailExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.UserDTOResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}
