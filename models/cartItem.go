package models

import "time"

type CartItem struct {
	ID         int     `gorm:"primarykey;not null;AUTO_INCREMENT" json:"id" form:"id"`
	CartID     int     `json:"cart_id" form:"cart_id"`
	Cart       Cart    `gorm:"foreignkey:CartID"`
	ProductID  int     `json:"product_id" form:"product_id"`
	Product    Product `gorm:"foreignkey:ProductID"`
	Quantity   int     `gorm:"type:int;not null" json:"quantity" form:"quantity"`
	AddedAt    time.Time
	CheckoutID int      `json:"checkout_id" form:"checkout_id"`
	Checkout   Checkout `gorm:"foreignkey:CheckoutID"`
}
