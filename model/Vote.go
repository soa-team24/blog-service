package model

import (
	"time"

	"github.com/google/uuid"
)

type Vote struct {
	ID           uuid.UUID `json:"id"`
	IsUpvote     bool      `json:"isUpvote"`
	UserID       uuid.UUID `json:"userId"`
	BlogId       uuid.UUID `json:"blogId"`
	CreationTime time.Time `json:"creationTime"`
}
