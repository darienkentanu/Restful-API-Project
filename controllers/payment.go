package controllers

import (
	"altastore/lib/database"
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
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid transaction ID")
	}
	status, err := payment.GetTransactionStatus(transactionID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "an error has been occured")
	}
	if status == 0 {
		return c.JSON(http.StatusOK, M{
			"status transaksi": "belum dibayar",
		})
	}
	return c.JSON(http.StatusOK, M{
		"status transaksi": "sudah dibayar",
	})
}
