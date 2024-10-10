package service

import (
	"errors"
	"fmt"
	"like-service/internal/model"
	"like-service/internal/repository"
	"net/http"
)

type LikeService interface {
	AddLike(userID, messageID int) (*model.Like, error)
	GetLikeCount(messageID int) (int64, error)
}

type likeService struct {
	repo repository.LikeRepository
}

func NewLikeService(repo repository.LikeRepository) LikeService {
	return &likeService{repo}
}

// ValidateUser checks with User Service if the user exists
func (s *likeService) ValidateUser(userID int) (bool, error) {
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

// ValidateMessage checks with Message Service if the message exists
func (s *likeService) ValidateMessage(messageID int) (bool, error) {
	url := fmt.Sprintf("http://message-service:8081/messages/%d", messageID)
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, errors.New("message not found")
}

func (s *likeService) AddLike(userID, messageID int) (*model.Like, error) {
	// Validate user
	validUser, err := s.ValidateUser(userID)
	if err != nil || !validUser {
		return nil, errors.New("invalid user")
	}

	// Validate message
	validMessage, err := s.ValidateMessage(messageID)
	if err != nil || !validMessage {
		return nil, errors.New("invalid message")
	}

	// Add like
	like := &model.Like{
		UserID:    userID,
		MessageID: messageID,
	}
	err = s.repo.AddLike(like)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (s *likeService) GetLikeCount(messageID int) (int64, error) {
	return s.repo.GetLikeCountByMessageID(messageID)
}
