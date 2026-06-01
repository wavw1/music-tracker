package idea_handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wavw1/music-tracker/internal/dto"
	core_errors "github.com/wavw1/music-tracker/internal/errors"
)

func (h *IdeaHandler) GetIdea(c *gin.Context) {
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

	idea, err := h.service.GetIdea(
		c.Request.Context(),
		ideaID,
		userID,
	)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
		ID:     idea.ID,
		Title:  idea.Title,
		Bpm:    idea.Bpm,
		Key:    idea.Key,
		Status: idea.Status,
		Tags:   idea.Tags,
	}

	c.JSON(http.StatusOK, response)
}
