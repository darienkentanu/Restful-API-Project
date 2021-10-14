package models

import "time"

type Product struct {
	ID            int        `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Name          string     `gorm:"type:varchar(255);unique;not null" json:"name" form:"name"`
	CategoryID	  int		 `json:"category_id" form:"category_id"`
	Category      Category	 `gorm:"foreignkey:CategoryID"`
	Description   string     `gorm:"type:varchar(255);not null" json:"description" form:"description"`
	Quantity      int        `gorm:"type:int;not null" json:"quantity" form:"quantity"`
	Price         int        `gorm:"type:int;not null" json:"price" form:"price"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}