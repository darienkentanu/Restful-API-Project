package models

import "time"

type Checkout struct {
	ID        int `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time
}

type CheckoutItems struct {
	ID       int
	CartID   int
	CartItem []CartItem
}
