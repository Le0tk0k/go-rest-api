package model

import "time"

type Category struct {
	ID        int       `gorm:"primary_key" json:"id`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
