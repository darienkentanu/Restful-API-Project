package models

import "time"

type Checkout struct {
	ID        int `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time
}

type CheckoutItems_Response struct {
	ID      int
	Product []CartItem_Response
}

type CheckoutItems_Input struct {
	Courier   string `json:"courier" form:"courier"`
	ProductID []int  `json:"product_id" form:"product_id"`
}

