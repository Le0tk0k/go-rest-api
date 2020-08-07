package model

import "time"

type Article struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Title     string    `json:"title" validate:"required"`
	Url       string    `json:"url" validate:"required"`
	Category  Category  `json:"category`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
