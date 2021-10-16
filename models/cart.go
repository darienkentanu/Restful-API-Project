package models

import "time"

type Cart struct {
	ID        int  `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	UserID    int  `gorm:"primarykey" json:"user_id" form:"user_id"`
	User      User `gorm:"foreignkey:UserID;" json:"-"`
	CreatedAt time.Time
}
