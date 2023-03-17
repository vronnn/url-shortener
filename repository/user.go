package repository

import (
	"context"
	"gin-gorm-clean-template/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user entity.User) (entity.User, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (entity.User, error)
	FindUserByID(ctx context.Context, userID uuid.UUID) (entity.User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) (error)
	UpdateUser(ctx context.Context, user entity.User) (error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func(db *userConnection) RegisterUser(ctx context.Context, user entity.User) (entity.User, error) {
	user.ID = uuid.New()
	tx := db.connection.Create(&user)
	if tx.Error != nil {
		return entity.User{}, tx.Error
	}
	return user, nil
}

func(db *userConnection) GetAllUser(ctx context.Context) ([]entity.User, error) {
	var listUser []entity.User
	tx := db.connection.Find(&listUser)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listUser, nil
}

func(db *userConnection) FindUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	tx := db.connection.Where("email = ?", email).Take(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func(db *userConnection) FindUserByID(ctx context.Context, userID uuid.UUID) (entity.User, error) {
	var user entity.User
	tx := db.connection.Where("id = ?", userID).Take(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func(db *userConnection) DeleteUser(ctx context.Context, userID uuid.UUID) (error) {
	tx := db.connection.Delete(&entity.User{}, &userID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func(db *userConnection) UpdateUser(ctx context.Context, user entity.User) (error) {
	tx := db.connection.Updates(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}