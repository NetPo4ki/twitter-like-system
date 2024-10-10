package repository

import (
	"errors"
	"message-service/internal/model"

	"gorm.io/gorm"
)

type MessageRepository interface {
	PostMessage(message *model.Message) error
	GetMessages(limit int) ([]model.Message, error)
	GetMessageByID(messageID int) (*model.Message, error)
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db}
}

func (r *messageRepository) PostMessage(message *model.Message) error {
	return r.db.Create(message).Error
}

func (r *messageRepository) GetMessages(limit int) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Order("created_at desc").Limit(limit).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *messageRepository) GetMessageByID(messageID int) (*model.Message, error) {
	var message model.Message
	err := r.db.Where("id = ?", messageID).First(&message, messageID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No message found for the given ID
		}
		return nil, err
	}
	return &message, nil
}
