package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wavw1/music-tracker/internal/db"
	idea_handler "github.com/wavw1/music-tracker/internal/idea/handler"
	idea_postgresql "github.com/wavw1/music-tracker/internal/idea/repository/postgreSQL"
	idea_service "github.com/wavw1/music-tracker/internal/idea/service"
	"github.com/wavw1/music-tracker/internal/middleware"
	user_handler "github.com/wavw1/music-tracker/internal/user/handler"
	user_postgresql "github.com/wavw1/music-tracker/internal/user/repository/postgreSQL"
	user_service "github.com/wavw1/music-tracker/internal/user/service"
)

func main() {
	router := gin.Default()
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := db.NewPostgresPool(ctx)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	ideaRepository := idea_postgresql.NewPostgreSql(pool)
	ideaService := idea_service.NewIdeaServiceStruct(ideaRepository)
	ideaHandler := idea_handler.NewIdeaHandler(ideaService)

	UserRepository := user_postgresql.NewPostgreSql(pool)
	userService := user_service.NewUserServiceStruct(UserRepository)
	userHandler := user_handler.NewUserHandler(userService)

	router.POST("/idea", middleware.AuthMiddleware(), ideaHandler.CreateIdea)
	router.GET("/idea", middleware.AuthMiddleware(), ideaHandler.GetIdeas)
	router.GET("/idea/:id", middleware.AuthMiddleware(), ideaHandler.GetIdea)
	router.PATCH("/idea/:id", middleware.AuthMiddleware(), ideaHandler.UpdateIdea)
	router.DELETE("/idea/:id", middleware.AuthMiddleware(), ideaHandler.DeleteIdea)

	router.POST("/auth/register", userHandler.Register)
	router.POST("/auth/login", userHandler.Login)

	router.Run()
}
