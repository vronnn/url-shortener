package dto

import (
	"github.com/google/uuid"
)

type CreateFollowingDTO struct {
	ID        	uuid.UUID   `gorm:"primary_key;not_null" json:"id"`

	UserID   	uuid.UUID 	`json:"user_id" form:"user_id"`
	FollowingID uuid.UUID 	`json:"following_id" form:"following_id" binding:"required"`
}