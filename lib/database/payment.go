package database

import (
	"altastore/config"
	"altastore/models"
	"errors"
)

func GetPaymentAmount(orderid int) (amount int, err error) {
	var t models.Transaction
	if err := config.InitDB().Where("order_id = ?", orderid).Last(&t).Error; err != nil {
		return 0, err
	}
	if t.Amount == 0 {
		return 0, errors.New("internal server error, -> payment amount couldn't be zero")
	}
	return t.Amount, nil
}
