package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status uint32

const (
	Draft     Status = 0
	Published Status = 1
	Closed    Status = 2
	Active    Status = 3
	Famous    Status = 4
)

type Category uint32

const (
	Destinations  Category = 0
	Travelogues   Category = 1
	Activities    Category = 2
	Gastronomy    Category = 3
	Tips          Category = 4
	Culture       Category = 5
	Accommodation Category = 6
)

type Blog struct {
	ID           uuid.UUID `json:"id"`
	UserID       uint32    `json:"userId"`
	Username     string    `json:"username"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreationTime time.Time `json:"creationTime"`
	Status       Status    `json:"status"`
	Image        string    `json:"image"`
	Category     Category  `json:"category"`
}

func (blog *Blog) BeforeCreate(scope *gorm.DB) error {
	blog.ID = uuid.New()
	return nil
}
