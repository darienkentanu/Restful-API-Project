package models

import "time"

type Transaction struct {
	ID            int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	OrderID       int    `gorm:"type:bigint" json:"order_id" form:"order_id"`
	UserID        int    `json:"user_id" form:"user_id"`
	User          User   `gorm:"foreignkey:UserID" json:"-"`
	Address       string `gorm:"type:longtext;" json:"address" form:"address"`
	Courier       string `gorm:"type:varchar(30);not null" json:"courier" form:"courier"`
	PaymentStatus int    `gorm:"type:tinyint;not null;default:0;" json:"payment_status" form:"payment_status"`
	Amount        int    `gorm:"type:int;not null" json:"amount" form:"amount"`
	CreatedAt     time.Time
	CheckoutID    int      `gorm:"primarykey;" json:"checkout_id" form:"checkout_id"`
	Checkout      Checkout `gorm:"foreignkey:CheckoutID;" json:"-"`
}

type TransactionReport struct {
	ID            int    `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	OrderID       int    `gorm:"type:bigint" json:"order_id" form:"order_id"`
	UserID        int    `json:"user_id" form:"user_id"`
	Address       string `gorm:"type:longtext;" json:"address" form:"address"`
	Courier       string `gorm:"type:varchar(30);not null" json:"courier" form:"courier"`
	PaymentStatus int    `gorm:"type:tinyint;not null;default:0;" json:"payment_status" form:"payment_status"`
	Amount        int    `gorm:"type:int;not null" json:"amount" form:"amount"`
	CreatedAt     string
	CheckoutID    int     `gorm:"primarykey;" json:"checkout_id" form:"checkout_id"`
}
