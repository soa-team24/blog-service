package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID               uuid.UUID `json:"id"`
	UserID           uint32    `json:"userId"`
	Username         string    `json:"username"`
	BlogId           uuid.UUID `json:"blogId"`
	Text             string    `json:"text"`
	CreationTime     time.Time `json:"creationTime"`
	LastModification time.Time `json:"lastModification"`
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	comment.ID = uuid.New()
	return nil
}
