package models

import (
	"time"
)

type User struct {
	ID                  int    		`gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	Fullname            string 		`gorm:"type:varchar(100);not null" json:"fullname" form:"fullname"`
	Username            string 		`gorm:"type:varchar(255);unique;not null" json:"username" form:"username"`
	Email               string 		`gorm:"type:varchar(100);unique;not null" json:"email" form:"email"`
	Password            string 		`gorm:"type:varchar(255);not null" json:"password" form:"password"`
	PhoneNumber			string		`gorm:"type:varchar(20);unique;not null" json:"phone_number" form:"phone_number"`
	Gender              string 		`gorm:"type:enum('male', 'female'); not null" json:"gender" form:"gender"`
	Address             string 		`gorm:"type:longtext;not null" json:"address" form:"address"`
	Role				string		`gorm:"type:enum('admin', 'customer');not null" json:"role" form:"role"`
	Token               string 		`gorm:"type:longtext;" json:"token" form:"token"`
	CartID				int			`gorm:"primarykey" json:"cart_id" form:"cart_id"`
	Cart	    		Cart        `gorm:"foreignkey:CartID;"`
	CreatedAt           time.Time
}