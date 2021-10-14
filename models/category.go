package models

import "time"

type Category struct {
	ID          int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Name        string `gorm:"type:varchar(255);unique;not null" json:"name" form:"id"`
	Description string `gorm:"type:varchar(255);not null" json:"description" form:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}