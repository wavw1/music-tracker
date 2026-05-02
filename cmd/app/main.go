package main //

import (
	"github.com/gin-gonic/gin"
	"github.com/wavw1/music-tracker/internal/handler"
)

func main() {
	router := gin.Default()

	httpHandlers := handler.NewHttpHandlers()

	router.POST("/idea", httpHandlers.CreateIdea)

	router.Run()
}
