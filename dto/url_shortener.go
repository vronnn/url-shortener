package dto

import (
	"github.com/google/uuid"
)

type UrlShortenerCreateDTO struct {
	ID        	uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	LongUrl 	string 		`json:"long_url" form:"long_url" binding:"required"`
	ShortUrl 	string 		`json:"short_url" form:"short_url" binding:"required"`
	Views 		uint64  	`json:"views" form:"views" binding:"required"`
	Private		bool		`json:"private" form:"private" binding:"required"`

	Password	string		`json:"password" form:"password"`
}

type UrlShortenerUpdateDTO struct {
	ID        	uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	LongUrl 	string 		`json:"long_url" form:"long_url"`
	ShortUrl 	string 		`json:"short_url" form:"short_url"`
	Views 		uint64  	`json:"views" form:"views"`
	Private		bool		`json:"private" form:"private"`

	Password	string		`json:"password" form:"password"`
}