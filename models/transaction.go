package models

import "time"

type Transaction struct {
	ID              int            `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	CustomerID		int			   `json:"customer_id"`
	Customer	    Customer       `gorm:"foreignkey:CustomerID"`
	Amount          int            `gorm:"type:int;not null" json:"amount"`
	CreatedAt       time.Time
}

type CheckoutItem struct {
	ID              int          `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	TransactionID	int			 `json:"transaction_id"`
	Transaction		Transaction  `gorm:"foreignkey:TransactionID"`
	ProductID		int			 `json:"product_id"`
	Product     	Product	     `gorm:"foreignkey:ProductID"`
	Quantity        int          `gorm:"type:bigint;not null" json:"quantity"`
	CreatedAt       time.Time
}