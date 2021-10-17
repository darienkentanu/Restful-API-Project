package models

type PaymentDetail struct {
	ID            int         `gorm:"primarykey;AUTO_INCREMENT" json:"id" form:"id"`
	PaymentMethod string      `gorm:"type:varchar(100);not null" json:"payment_method" form:"payment_method"`
	TransactionID int         `gorm:"not null" json:"transaction_id" form:"transaction_id"`
	Transaction   Transaction `gorm:"foreignkey:TransactionID" json:"-"`
	Amount        int         `gorm:"type:int;not null" json:"amount" form:"amount"`
}
