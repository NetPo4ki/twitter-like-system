package main

import (
	"like-service/internal/db"
	"like-service/internal/handler"
	"like-service/internal/repository"
	"like-service/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbConn, err := db.Connect()
	if err != nil {
		panic(err)
	}

	likeRepo := repository.NewLikeRepository(dbConn)
	likeService := service.NewLikeService(likeRepo)
	likeHandler := handler.NewLikeHandler(likeService)

	r := gin.Default()

	r.POST("/likes", likeHandler.AddLike)
	r.GET("/messages/:messageID/likes", likeHandler.GetLikeCount)

	r.Run(":8082") // Runs on port 8082
}
