package idea_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wavw1/music-tracker/internal/dto"
)

func (h *IdeaHandler) GetIdeas(c *gin.Context) {
	var limitParam = c.DefaultQuery("limit", "0")
	limit, err := strconv.Atoi(limitParam)
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

	var offsetParam = c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ideas, err := h.service.GetIdeas(
		c.Request.Context(),
		limit,
		offset,
		userID,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": fmt.Sprintf("failed to handle service. error: %v", err)},
		)
		return
	}

	var ideasResponse = []dto.IdeaDTOResponse{}

	for _, idea := range ideas {
		ideaResponse := dto.IdeaDTOResponse{
			ID:     idea.ID,
			Title:  idea.Title,
			Bpm:    idea.Bpm,
			Key:    idea.Key,
			Status: idea.Status,
			Tags:   idea.Tags,
		}

		ideasResponse = append(ideasResponse, ideaResponse)
	}

	var response = dto.IdeasDTOResponse{
		Ideas: ideasResponse,
	}

	c.JSON(http.StatusOK, response)
}
