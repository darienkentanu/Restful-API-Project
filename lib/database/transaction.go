package database

import (
	"altastore/config"
	"altastore/models"
)

func UpdateTransactionStatus(TransactionID int) error {
	var t models.Transaction
	if err := config.InitDB().First(&t, TransactionID).Error; err != nil {
		return err
	}
	t.PaymentStatus = 1
	if err := config.InitDB().Save(&t).Error; err != nil {
		return err
	}
	return nil
}
