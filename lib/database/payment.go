package database

import (
	"altastore/config"
	"altastore/models"
)

func GetPaymentAmount(idTransaction int) (amount int, err error) {
	var p models.PaymentDetail
	if err := config.InitDB().Where("transaction_id = ?", idTransaction).Last(&p).Error; err != nil {
		return 0, err
	}
	return p.Amount, nil
}
