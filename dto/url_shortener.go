package dto

import (
	"gin-gorm-clean-template/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UrlShortenerCreateDTO struct {
	ID        	uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	Title 		string 		`json:"title" form:"title" binding:"required"`
	LongUrl 	string 		`json:"long_url" form:"long_url" binding:"required"`
	ShortUrl 	string 		`json:"short_url" form:"short_url" binding:"required"`
	Views 		uint64  	`json:"views" form:"views"`
	IsPrivate	*bool		`json:"is_private" form:"is_private" binding:"required"`
	IsFeeds		*bool		`json:"is_feeds" form:"is_feeds" binding:"required"`

	UserID   	uuid.UUID 		`gorm:"foreignKey" json:"user_id"`
	User     	*entity.User  	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`

	Password	string		`json:"password" form:"password"`
}

type UrlShortenerUpdateDTO struct {
	ID        	uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	Title 		string 		`json:"title" form:"title"`
	LongUrl 	string 		`json:"long_url" form:"long_url"`
	ShortUrl 	string 		`json:"short_url" form:"short_url"`
	Views 		uint64  	`json:"views" form:"views"`

	Password	string		`json:"password" form:"password"`
}

type UrlShortenerResponseDTO struct {
	ID        	uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	Title 		string 		`json:"title" form:"title"`
	LongUrl 	string 		`json:"long_url" form:"long_url"`
	ShortUrl 	string 		`json:"short_url" form:"short_url"`
	Views 		uint64  	`json:"views" form:"views"`
	IsPrivate	*bool		`json:"is_private" form:"is_private"`
	IsFeeds		*bool		`json:"is_feeds" form:"is_feeds"`
	Username	string		`json:"username"`

	UserID   	uuid.UUID 		`gorm:"foreignKey" json:"user_id"`

	Password	string		`json:"password" form:"password"`
	CreatedAt 		time.Time 	`json:"created_at" default:"CURRENT_TIMESTAMP"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
	DeletedAt 		gorm.DeletedAt
}

func BoolPointer(b bool) *bool {
    return &b
}