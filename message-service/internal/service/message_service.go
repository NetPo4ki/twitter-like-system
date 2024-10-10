package service

import (
	"errors"
	"fmt"
	"message-service/internal/model"
	"message-service/internal/repository"
	"net/http"
)

type MessageService interface {
	PostMessage(userID int, content string) (*model.Message, error)
	GetMessages(limit int) ([]model.Message, error)
	GetMessageByID(messageID int) (*model.Message, error)
}

type messageService struct {
	repo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) MessageService {
	return &messageService{repo}
}

// ValidateUser checks with User Service if the user exists
func (s *messageService) ValidateUser(userID int) (bool, error) {
	url := fmt.Sprintf("http://user-service:8080/users/%d", userID)
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, errors.New("user not found")
}

func (s *messageService) PostMessage(userID int, content string) (*model.Message, error) {
	// Validate user
	validUser, err := s.ValidateUser(userID)
	if err != nil || !validUser {
		return nil, errors.New("invalid user")
	}

	// Post the message
	message := &model.Message{
		UserID:  userID,
		Content: content,
	}
	err = s.repo.PostMessage(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (s *messageService) GetMessages(limit int) ([]model.Message, error) {
	return s.repo.GetMessages(limit)
}

func (s *messageService) GetMessageByID(messageID int) (*model.Message, error) {
	message, err := s.repo.GetMessageByID(messageID)

	if err != nil {
		return nil, err
	}

	if message == nil {
		return nil, errors.New("message not found")
	}

	return message, nil
}
