package models

import "time"

type Product struct {
	ID            int        `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Name          string     `gorm:"type:varchar(255);unique;not null" json:"name"`
	CategoryID	  int		 `json:"category_id"`
	Category      Category	 `gorm:"foreignkey:CategoryID"`
	Description   string     `gorm:"type:varchar(255);not null" json:"description"`
	Quantity      int        `gorm:"type:bigint;not null" json:"quantity"`
	Price         int        `gorm:"type:int;not null" json:"price"`
	Unit          string     `gorm:"type:varchar(255);not null" json:"unit"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}