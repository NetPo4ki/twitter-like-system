package main

import (
	"log"
	"user-service/internal/db"
	"user-service/internal/handler"
	"user-service/internal/repository"
	"user-service/internal/service"

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

	userRepo := repository.NewUserRepository(dbConn)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	r.POST("/users", userHandler.RegisterUser)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.GET("/users", userHandler.ListUsers)

	r.Run(":8080")
}
