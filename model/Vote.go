package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vote struct {
	ID           uuid.UUID `json:"id"`
	IsUpvote     bool      `json:"isUpvote"`
	UserID       uint32    `json:"userId"`
	BlogId       uuid.UUID `json:"blogId"`
	CreationTime time.Time `json:"creationTime"`
}

func (vote *Vote) BeforeCreate(scope *gorm.DB) error {
	vote.ID = uuid.New()
	return nil
}
