package models

import "time"

type Category struct {
	ID          int    `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Name        string `gorm:"type:varchar(255);unique;not null" json:"name"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}