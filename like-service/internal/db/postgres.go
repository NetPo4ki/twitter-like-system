package db

import (
	"fmt"
	"like-service/internal/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("LIKE_DB_NAME")
	host := os.Getenv("POSTGRES_HOST_LIKE")
	port := os.Getenv("POSTGRES_PORT_LIKE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Auto-migrate the Like schema
	err = db.AutoMigrate(&model.Like{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	return db, nil
}
