package model

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID               uuid.UUID `json:"id"`
	UserID           uuid.UUID `json:"authorId"`
	Username         string    `json:"authorUsername"`
	BlogId           uuid.UUID `json:"blogId"`
	Text             string    `json:"text"`
	CreationTime     time.Time `json:"creationTime"`
	LastModification time.Time `json:"lastModification"`
}
