package db

import (
	"fmt"
	"log"
	"message-service/internal/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("MESSAGE_DB_NAME")
	host := os.Getenv("POSTGRES_HOST_MESSAGE")
	port := os.Getenv("POSTGRES_PORT_MESSAGE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Auto-migrate the Message schema
	err = db.AutoMigrate(&model.Message{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	return db, nil
}
