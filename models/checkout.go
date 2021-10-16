package models

import "time"

type Checkout struct {
	ID        int `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time
}
