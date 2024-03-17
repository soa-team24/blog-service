package model

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Draft     Status = "draft"
	Published Status = "published"
	Closed    Status = "closed"
	Active    Status = "active"
	Famous    Status = "famous"
)

type Category string

const (
	Destinations  Category = "destinations"
	Travelogues   Category = "travelogues"
	Activities    Category = "activities"
	Gastronomy    Category = "gastronomy"
	Tips          Category = "tips"
	Culture       Category = "culture"
	Accommodation Category = "accommodation"
)

type Blog struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"authorId"`
	Username     string    `json:"authorUsername"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreationTime time.Time `json:"creationTime"`
	Status       Status    `json:"status"`
	Image        string    `json:"image"`
	Category     Category  `json:"category"`
}
