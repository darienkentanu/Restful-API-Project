package models

import (
	"time"
)

type Customer struct {
	ID                  int    		`gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Fullname            string 		`gorm:"type:varchar(100);not null" json:"fullname"`
	Username            string 		`gorm:"type:varchar(255);unique;not null" json:"username"`
	Email               string 		`gorm:"type:varchar(100);unique;not null" json:"email"`
	Password            string 		`gorm:"type:varchar(255);not null" json:"password"`
	Gender              string 		`gorm:"type:varchar(30);not null" json:"gender"`
	Address             string 		`gorm:"type:longtext;" json:"address"`
	BankName           	string 		`gorm:"type:varchar(255);" json:"bank_name"`
	BankAccountNumber 	int    		`gorm:"type:bigint;default:0;" json:"bank_account_number"`
	Token               string 		`gorm:"type:longtext;" json:"token"`
	CreatedAt           time.Time
}