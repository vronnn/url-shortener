package service

import (
	"context"
	"gin-gorm-clean-template/dto"
	"gin-gorm-clean-template/entity"
	"gin-gorm-clean-template/repository"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type FollowingService interface {
	CreateFollowing(ctx context.Context, followingDTO dto.CreateFollowingDTO) (entity.Following, error)
	FindFollowingByUserID(ctx context.Context, userID string) ([]entity.Following, error)
	CheckDuplicate(ctx context.Context, userID string, followingID string) (bool)
}

type followingService struct {
	followingRepository repository.FollowingRepository
}

func NewFollowingService(fr repository.FollowingRepository) FollowingService {
	return &followingService{
		followingRepository: fr,
	}
}

func(fs *followingService) CreateFollowing(ctx context.Context, followingDTO dto.CreateFollowingDTO) (entity.Following, error) {
	following := entity.Following{}
	err := smapping.FillStruct(&following, smapping.MapFields(followingDTO))
	if err != nil {
		return following, err
	}
	return fs.followingRepository.CreateFollowing(ctx, following)
}

func(fs *followingService) FindFollowingByUserID(ctx context.Context, userID string) ([]entity.Following, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	return fs.followingRepository.FindFollowingByUserID(ctx, userUUID)
}

func(fs *followingService) CheckDuplicate(ctx context.Context, userID string, followingID string) (bool) {
	userUUID, _ := uuid.Parse(userID)
	followingUUID, _ := uuid.Parse(followingID)
	return fs.followingRepository.CheckDuplicate(ctx, userUUID, followingUUID)
}