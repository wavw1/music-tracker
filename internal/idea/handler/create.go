package idea_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wavw1/music-tracker/internal/dto"
	"github.com/wavw1/music-tracker/internal/model"
)

func (h *IdeaHandler) CreateIdea(
	c *gin.Context,
) {
	var ideaRequest dto.IdeaDTORequest

	getUserID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get user_id"})
		return
	}
	userID := getUserID.(int64)

	if err := c.ShouldBindJSON(&ideaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ideaModel = model.Idea{
		Title: ideaRequest.Title,
		Bpm:   ideaRequest.Bpm,
		Key:   ideaRequest.Key,
		Tags:  ideaRequest.Tags,
	}

	ideaResponse, err := h.service.CreateIdea(
		c.Request.Context(),
		ideaModel,
		userID,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := dto.IdeaDTOResponse{
		ID:     ideaResponse.ID,
		Title:  ideaResponse.Title,
		Bpm:    ideaResponse.Bpm,
		Key:    ideaResponse.Key,
		Status: ideaResponse.Status,
		Tags:   ideaResponse.Tags,
	}

	c.JSON(http.StatusCreated, response)
}
