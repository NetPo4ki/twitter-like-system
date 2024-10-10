package main

import (
	"log"
	"message-service/internal/db"
	"message-service/internal/handler"
	"message-service/internal/repository"
	"message-service/internal/service"

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

	messageRepo := repository.NewMessageRepository(dbConn)
	messageService := service.NewMessageService(messageRepo)
	messageHandler := handler.NewMessageHandler(messageService)

	r := gin.Default()

	r.POST("/messages", messageHandler.PostMessage)
	r.GET("/messages", messageHandler.GetMessages)
	r.GET("/messages/:id", messageHandler.GetMessageByID)

	r.Run(":8081") // Runs on port 8081
}
