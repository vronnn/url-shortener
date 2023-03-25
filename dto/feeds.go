package dto

import "github.com/google/uuid"

type FeedsResponseDTO struct {
	ID        	uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	Data 		string 		`json:"data"`
	Method		string		`json:"method"`

	UrlShortenerID   	uuid.UUID 			`gorm:"foreignKey" json:"url_shortener_id"`
}