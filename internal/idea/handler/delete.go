package idea_handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	core_errors "github.com/wavw1/music-tracker/internal/errors"
)

func (h *IdeaHandler) DeleteIdea(c *gin.Context) {
	idParam := c.Param("id")
	ideaID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	getUserID, ok := c.Get("user_id")
	fmt.Println(getUserID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get user_id"})
		return
	}
	userID := getUserID.(int64)

	if err := h.service.DeleteIdea(
		c.Request.Context(),
		ideaID,
		userID,
	); err != nil {
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

	c.Status(http.StatusNoContent)
}
