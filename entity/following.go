package entity

import (
	"github.com/google/uuid"
)

type Following struct {
	ID        	uuid.UUID   `gorm:"primary_key;not_null" json:"id"`

	UserID   	uuid.UUID 	`gorm:"foreignKey" json:"user_id"`
	FollowingID uuid.UUID 	`gorm:"foreignKey" json:"following_id"`
	User     	*User  		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	
	Timestamp
}