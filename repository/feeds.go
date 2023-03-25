package repository

import (
	"context"
	"gin-gorm-clean-template/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FeedsRepository interface {
	CreateFeeds(ctx context.Context, feeds entity.Feeds) (entity.Feeds, error)
	GetAllFeeds(ctx context.Context) ([]entity.Feeds, error)
}

type feedsConnection struct {
	connection *gorm.DB
}

func NewFeedsRepository(db *gorm.DB) FeedsRepository {
	return &feedsConnection{
		connection: db,
	}
}

func(db *feedsConnection) CreateFeeds(ctx context.Context, feeds entity.Feeds) (entity.Feeds, error) {
	feeds.ID = uuid.New()
	tx := db.connection.Create(&feeds)
	if tx.Error != nil {
		return entity.Feeds{}, tx.Error
	}
	return feeds, nil
}

func(db *feedsConnection) GetAllFeeds(ctx context.Context) ([]entity.Feeds, error) {
	var feedsList []entity.Feeds
	tx := db.connection.Find(&feedsList)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return feedsList, nil
}