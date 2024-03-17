package model

import "github.com/google/uuid"

type Comment struct {
	ID uuid.UUID `json:"id"`
}
