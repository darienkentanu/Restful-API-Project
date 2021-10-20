package models

import "time"

type CartItem struct {
	ID         int     `gorm:"primarykey;not null;AUTO_INCREMENT" json:"id" form:"id"`
	CartID     int     `json:"cart_id" form:"cart_id"`
	Cart       Cart    `gorm:"foreignkey:CartID" json:"-"`
	ProductID  int     `json:"product_id" form:"product_id"`
	Product    Product `gorm:"foreignkey:ProductID" json:"-"`
	Quantity   int     `gorm:"type:int;not null" json:"quantity" form:"quantity"`
	CreatedAt  time.Time
	CheckoutID int      `json:"-"`
	Checkout   Checkout `gorm:"foreignkey:CheckoutID;null" json:"-"`
}

type AddCartItem struct {
	ProductID int `json:"product_id" form:"product_id"`
	Quantity  int `json:"quantity" form:"quantity"`
}

type UpdateCartItem struct {
	Quantity int `json:"quantity" form:"quantity"`
}

type CartItem_Response struct {
	ProductID    int    `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPrice int    `json:"product_price"`
	Quantity     int    `json:"quantity"`
}
