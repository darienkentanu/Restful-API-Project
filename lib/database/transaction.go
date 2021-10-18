package database

import (
	"altastore/config"
	"altastore/models"
	"net/http"

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
