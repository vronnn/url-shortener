package service

import (
	"context"
	"gin-gorm-clean-template/dto"
	"gin-gorm-clean-template/entity"
	"gin-gorm-clean-template/repository"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type UrlShortenerService interface {
	CreateUrlShortener(ctx context.Context, urlShortenerDTO dto.UrlShortenerCreateDTO) (entity.UrlShortener, error)
	GetUrlShortenerByShortUrl(ctx context.Context, shortUrl string) (entity.UrlShortener, error)
}

type urlShortenerService struct {
	urlShortenerRepository repository.UrlShortenerRepository
	privateRepository repository.PrivateRepository
}

func NewUrlShortenerService(ur repository.UrlShortenerRepository, pr repository.PrivateRepository) UrlShortenerService {
	return &urlShortenerService{
		urlShortenerRepository: ur,
		privateRepository: pr,
	}
}

func(us *urlShortenerService) CreateUrlShortener(ctx context.Context, urlShortenerDTO dto.UrlShortenerCreateDTO) (entity.UrlShortener, error) {
	urlShortener := entity.UrlShortener{}
	err := smapping.FillStruct(&urlShortener, smapping.MapFields(urlShortenerDTO))
	if err != nil {
		return urlShortener, err
	}
	if *urlShortener.UserID == uuid.Nil {
		urlShortener.UserID = nil
	}
	res, err := us.urlShortenerRepository.CreateUrlShortener(ctx, urlShortener)
	if err != nil {
		return urlShortener, err
	}
	if *urlShortenerDTO.Private {
		private := entity.Private{
			Password: urlShortenerDTO.Password,
			UrlShortenerID: res.ID,
		}
		_, err = us.privateRepository.CreatePrivate(ctx, private)
		if err != nil {
			return urlShortener, err
		}
	}
	return res, err
}

func(us *urlShortenerService) GetUrlShortenerByShortUrl(ctx context.Context, shortUrl string) (entity.UrlShortener, error) {
	return us.urlShortenerRepository.GetUrlShortenerByShortUrl(ctx, shortUrl)
}