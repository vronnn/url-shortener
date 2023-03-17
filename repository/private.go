package repository

import (
	"context"
	"gin-gorm-clean-template/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PrivateRepository interface {
	CreatePrivate(ctx context.Context, private entity.Private) (entity.Private, error)
	GetPrivateByUrlShortenerID(ctx context.Context, urlShortenerID uuid.UUID) (entity.Private, error)
	UpdatePrivate(ctx context.Context, private entity.Private) (error)
}

type privateConnection struct {
	connection *gorm.DB
}

func NewPrivateRepository(db *gorm.DB) PrivateRepository {
	return &privateConnection{
		connection: db,
	}
}

func(db *privateConnection) CreatePrivate(ctx context.Context, private entity.Private) (entity.Private, error) {
	private.ID = uuid.New()
	tx := db.connection.Create(&private)
	if tx.Error != nil {
		return entity.Private{}, tx.Error
	}
	return private, nil
}

func(db *privateConnection) GetPrivateByUrlShortenerID(ctx context.Context, urlShortenerID uuid.UUID) (entity.Private, error) {
	var private entity.Private
	tx := db.connection.Where("id = ?", urlShortenerID).Take(&private)
	if tx.Error != nil {
		return entity.Private{}, tx.Error
	}
	return private, nil
}

func(db *privateConnection) UpdatePrivate(ctx context.Context, private entity.Private) (error) {
	tx := db.connection.Updates(&private)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}