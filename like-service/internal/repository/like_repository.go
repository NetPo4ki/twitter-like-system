package repository

import (
	"like-service/internal/model"

	"gorm.io/gorm"
)

type LikeRepository interface {
	AddLike(like *model.Like) error
	GetLikeCountByMessageID(messageID int) (int64, error)
}

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &likeRepository{db}
}

func (r *likeRepository) AddLike(like *model.Like) error {
	return r.db.Create(like).Error
}

func (r *likeRepository) GetLikeCountByMessageID(messageID int) (int64, error) {
	var count int64
	err := r.db.Model(&model.Like{}).Where("message_id = ?", messageID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
