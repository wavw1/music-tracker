package idea_handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wavw1/music-tracker/internal/dto"
	core_errors "github.com/wavw1/music-tracker/internal/errors"
	idea_domain "github.com/wavw1/music-tracker/internal/idea/domain"
)

type IdeaDTOUpdateRequest struct {
	Title *string `json:"title,omitempty"`
	Bpm   *int    `json:"bpm,omitempty"`
	Key   *string `json:"key,omitempty"`
}

func (h *IdeaHandler) UpdateIdea(c *gin.Context) {
	idParam := c.Param("id")
	ideaID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	getUserID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get user_id"})
		return
	}
	userID := getUserID.(int64)

	var ideaRequest IdeaDTOUpdateRequest

	if err := c.ShouldBindJSON(&ideaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ideaService = idea_domain.IdeaServiceInput{
		Title: ideaRequest.Title,
		Bpm:   ideaRequest.Bpm,
		Key:   ideaRequest.Key,
	}

	updatedIdea, err := h.service.UpdateIdea(
		c.Request.Context(),
		ideaID,
		ideaService,
		userID,
	)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if errors.Is(err, core_errors.ErrInvalidBPM) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if errors.Is(err, core_errors.ErrForbidden) {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.IdeaDTOResponse{
		ID:     updatedIdea.ID,
		Title:  updatedIdea.Title,
		Bpm:    updatedIdea.Bpm,
		Key:    updatedIdea.Key,
		Status: updatedIdea.Status,
		Tags:   updatedIdea.Tags,
	}

	c.JSON(http.StatusOK, response)
}
