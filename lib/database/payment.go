package database

import (
	"altastore/config"
	"altastore/models"
	"errors"
)

func GetPaymentAmount(idTransaction int) (amount int, err error) {
	var p models.PaymentDetail
	if err := config.InitDB().Where("transaction_id = ?", idTransaction).Last(&p).Error; err != nil {
		return 0, err
	}
	if p.Amount == 0 {
		return 0, errors.New("internal server error, -> payment amount couldn't be zero")
	}
	return p.Amount, nil
}
