package controllers

import (
	"altastore/lib/database"
	"altastore/models"
	"altastore/payment"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RequestBilling(c echo.Context) error {
	id := c.Param("id")
	idStr, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid transaction id")
	}
	amount, err := database.GetPaymentAmount(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "an error has been occured")
	}
	redirectURL, err := payment.RequestBilling(idStr, amount)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "an error has been occured")
	}
	return c.JSON(http.StatusOK, redirectURL)
}

func GetTransactionStatus(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid transaction ID")
	}
	status, err := payment.GetTransactionStatus(orderID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "an error has been occured")
	}
	if status == 0 {
		return c.JSON(http.StatusOK, M{
			"status transaksi": "belum dibayar",
		})
	}

	var paymentDetail models.PaymentDetail
	if status == 1 {
		transaction, err := database.GetTransactionByOrderID(orderID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		err = database.UpdateTransactionStatus(transaction.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "an error has been occured")
		}

		row := database.GetRowPaymentDetail(transaction.ID)
		if row == 0 {
			
			paymentDetail.PaymentMethod = "midtrans gateway"
			paymentDetail.TransactionID = transaction.ID

			paymentDetail, err = database.CreatePaymentDetail(paymentDetail)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
			}

			
			checkoutID := transaction.CheckoutID
			cartItems, err := database.GetCartItemByCheckoutID(checkoutID)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
			}

			for _, item := range cartItems {
				productID := item.ProductID
				product, err := database.GetProduct(productID)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
				}
				newQuantity := product.Quantity - item.Quantity
				product, err = database.UpdateProductQuantity(productID, newQuantity)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
				}
			}
		} else {
			paymentDetail, err = database.GetPaymentDetailByTransactionID(transaction.ID)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
			}
		}
	}
	return c.JSON(http.StatusOK, M{
		"status transaksi": "sudah dibayar",
		"data": paymentDetail,
	})
}
