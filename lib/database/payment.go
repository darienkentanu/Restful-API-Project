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

func GetRowPaymentDetail(transactionID int) int {
	var paymentDetail models.PaymentDetail
	row := config.InitDB().Where("transaction_id = ?", transactionID).Find(&paymentDetail).RowsAffected
	return int(row)
}

func CreatePaymentDetail(paymentDetail models.PaymentDetail) (models.PaymentDetail, error) {
	if err := config.InitDB().Save(&paymentDetail).Error; err != nil {
		return paymentDetail, err
	}

	return paymentDetail, nil
}

func GetPaymentDetailByTransactionID(transactionID int) (models.PaymentDetail, error) {
	var paymentDetail models.PaymentDetail

	if err := config.InitDB().Where("transaction_id = ?", transactionID).Find(&paymentDetail).Error; err != nil {
		return paymentDetail, err
	}

	return paymentDetail, nil
}