package models

import (
	"time"
)

type Customer struct {
	ID                  int    		`gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Fullname            string 		`gorm:"type:varchar(100);not null" json:"fullname" form:"fullname"`
	Username            string 		`gorm:"type:varchar(255);unique;not null" json:"username" form:"username"`
	Email               string 		`gorm:"type:varchar(100);unique;not null" json:"email" form:"email"`
	Password            string 		`gorm:"type:varchar(255);not null" json:"password" form:"password"`
	Gender              string 		`gorm:"type:varchar(30);not null" json:"gender" form:"gender"`
	Address             string 		`gorm:"type:longtext;" json:"address" form:"address"`
	Token               string 		`gorm:"type:longtext;" json:"token" form:"token"`
	CartID				int			`gorm:"primarykey" json:"cart_id" form:"cart_id"`
	Cart	    		Cart        `gorm:"foreignkey:CartID;"`
	CreatedAt           time.Time
}