package repository

import (
	"context"
	"gin-gorm-clean-template/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FollowingRepository interface {
	CreateFollowing(ctx context.Context, following entity.Following) (entity.Following, error)
	FindFollowingByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Following, error)
	CheckDuplicate(ctx context.Context, userID uuid.UUID, followingID uuid.UUID) (bool)
}

type followingConnection struct {
	connection *gorm.DB
}

func NewFollowingRepository(db *gorm.DB) FollowingRepository {
	return &followingConnection{
		connection: db,
	}
}

func(db *followingConnection) CreateFollowing(ctx context.Context, following entity.Following) (entity.Following, error) {
	following.ID = uuid.New()
	tx := db.connection.Create(&following)
	if tx.Error != nil {
		return entity.Following{}, tx.Error
	}
	return following, nil
}

func(db *followingConnection) FindFollowingByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Following, error) {
	var following []entity.Following
	tx := db.connection.Where("user_id = ?", userID).Find(&following)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return following, nil
}

func(db *followingConnection) CheckDuplicate(ctx context.Context, userID uuid.UUID, followingID uuid.UUID) (bool) {
	var following entity.Following
	tx := db.connection.Where("user_id = ? AND following_id = ?", userID, followingID).Take(&following)
	if tx.Error != nil {
		return true
	}
	return false
}