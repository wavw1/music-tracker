package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandlers) CreateIdea(c *gin.Context) {
	c.Status(http.StatusOK)
}
