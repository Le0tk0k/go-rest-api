package domain

import "time"

type Article struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	Category  *Category `json:"category"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
