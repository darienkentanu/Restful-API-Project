package models

import "time"

type Checkout struct {
	ID              int          `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	CreatedAt       time.Time
}

type Transaction struct {
	ID              int            `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	CustomerID		int			   `json:"customer_id" form:"customer_id"`
	Customer	    Customer       `gorm:"foreignkey:CustomerID"`
	Address         string 		   `gorm:"type:longtext;" json:"address" form:"address"`
	Courier			string 		   `gorm:"type:varchar(30);not null" json:"courier" form:"courier"`
	PaymentStatus	int			   `gorm:"type:tinyint;not null;default:0;" json:"payment_status" form:"payment_status"`
	CreatedAt       time.Time
	CheckoutID		int			   `gorm:"primarykey;" json:"checkout_id" form:"checkout_id"`
	Checkout	    Checkout       `gorm:"foreignkey:CheckoutID;"`
}