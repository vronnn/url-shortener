package service

import (
	"context"
	"gin-gorm-clean-template/entity"
	"gin-gorm-clean-template/repository"
)

type FeedsService interface {
	GetAllFeeds(ctx context.Context) ([]entity.Feeds, error)
}

type feedsService struct {
	feedsRepository repository.FeedsRepository
}

func NewFeedsService(fr repository.FeedsRepository) FeedsService {
	return &feedsService{
		feedsRepository: fr,
	}
}

func(fs *feedsService) GetAllFeeds(ctx context.Context) ([]entity.Feeds, error) {
	return fs.feedsRepository.GetAllFeeds(ctx)
}