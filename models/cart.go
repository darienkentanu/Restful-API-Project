package models

import "time"

type Cart struct {
	ID           int         `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	CustomerID	 int		 `json:"customer_id"`
	Customer     Customer	 `gorm:"foreignkey:CustomerID"`
	CreatedAt    time.Time
}

type CartItem struct {
	ID          int      `gorm:"primarykey;not null;AUTO_INCREMENT" json:"id"`
	CartID		int		 `json:"cart_id"`
	Cart     	Cart	 `gorm:"foreignkey:CartID"`
	ProductID	int		 `json:"product_id"`
	Product     Product	 `gorm:"foreignkey:ProductID"`
	Quantity    int      `gorm:"type:bigint;not null" json:"quantity"`
}

