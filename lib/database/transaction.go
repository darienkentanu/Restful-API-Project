package database

import (
	"altastore/config"
	"altastore/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
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

func AddTransaction(courier string, orderID, amount, userID, checkoutID int) error {
	var t models.Transaction
	t.Courier = courier
	t.OrderID = orderID
	t.Amount = amount
	user, err := GetUserByID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	t.Address = user.Address
	t.UserID = user.ID
	t.CheckoutID = checkoutID
	if err := config.InitDB().Save(&t).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot save data to transaction table")
	}
	return nil
}

func GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction

	
	if err := config.InitDB().Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func GetTransactionsByUserID(userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction

	if err := config.InitDB().Find(&transactions, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func GetTransationsRangeDate(rangeDate string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	today := time.Now()
	lastWeek := today; // today - 7 days
	lastMonth := today; // today - 30 days

	if rangeDate == "daily" {
		if err :=  config.InitDB().Where("created_at = ?", today).Find(&transactions).Error; err != nil {
			return nil, err
		}
	} else if rangeDate == "weekly" {
		if err :=  config.InitDB().Where("created_at >= ?", lastWeek).Find(&transactions).Error; err != nil {
			return nil, err
		}
	} else if rangeDate == "monthly" {
		if err :=  config.InitDB().Where("created_at >= ?", lastMonth).Find(&transactions).Error; err != nil {
			return nil, err
		}
	}

	return transactions, nil
}