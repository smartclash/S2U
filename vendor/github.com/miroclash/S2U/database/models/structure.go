package models

import "time"

type Structure struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"index;column:deleted_at"`
}
