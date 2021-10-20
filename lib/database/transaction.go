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

func GetTransationsRangeDate(rangeDate string) ([]models.TransactionReport, error) {
	var transactions []models.TransactionReport

	if rangeDate == "daily" {
		rows, err := config.InitDBSQL().Query("SELECT * FROM transactions WHERE created_at >= DATE_ADD(CURDATE(), INTERVAL -1 DAY)")
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		
		for rows.Next() {
        	var trans models.TransactionReport
			if err := rows.Scan(&trans.ID, &trans.OrderID, &trans.UserID,
				&trans.Address, &trans.Courier, &trans.PaymentStatus, &trans.Amount, &trans.CreatedAt, &trans.CheckoutID); err != nil {
				return nil, err
			}
			transactions = append(transactions, trans)
		}
	} else if rangeDate == "weekly" {
		rows, err := config.InitDBSQL().Query("SELECT * FROM transactions WHERE created_at >= DATE_ADD(CURDATE(), INTERVAL -7 DAY)")
		if err != nil {
			return nil, err
		}
		
		for rows.Next() {
        	var trans models.TransactionReport
			if err := rows.Scan(&trans.ID, &trans.OrderID, &trans.UserID,
				&trans.Address, &trans.Courier, &trans.PaymentStatus, &trans.Amount, &trans.CreatedAt, &trans.CheckoutID); err != nil {
				return nil, err
			}
			transactions = append(transactions, trans)
		}
	} else if rangeDate == "monthly" {
		rows, err := config.InitDBSQL().Query("SELECT * FROM transactions WHERE created_at >= DATE_ADD(CURDATE(), INTERVAL -30 DAY)")
		if err != nil {
			return nil, err
		}
		
		for rows.Next() {
        	var trans models.TransactionReport
			if err := rows.Scan(&trans.ID, &trans.OrderID, &trans.UserID,
				&trans.Address, &trans.Courier, &trans.PaymentStatus, &trans.Amount, &trans.CreatedAt, &trans.CheckoutID); err != nil {
				return nil, err
			}
			transactions = append(transactions, trans)
		}
	}

	return transactions, nil
}

func GetTransactionByOrderID(orderID int) (models.Transaction, error) {
	var transaction models.Transaction

	if err := config.InitDB().Find(&transaction, "order_id = ?", orderID).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func GetCartItemByCheckoutID(checkoutID int) ([]models.CartItem, error) {
	var cartItems []models.CartItem

	if err := config.InitDB().Find(&cartItems, "checkout_id = ?", checkoutID).Error; err != nil {
		return nil, err
	}

	return cartItems, nil
}