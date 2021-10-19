package controllers

import (
	"altastore/lib/database"
	"altastore/middlewares"
	"altastore/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllTransactionsController(c echo.Context) error {
	var transactions []models.Transaction
	roleUser := middlewares.CurrentRoleLoginUser(c)
	userID := middlewares.CurrentLoginUser(c)

	var err error
	if roleUser == "admin" {
		transactions, err = database.GetAllTransactions()

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
	} else if roleUser == "customer" {
		transactions, err = database.GetTransactionsByUserID(userID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   transactions,
	})
}

func GetTransactionsWithRangeDate(c echo.Context) error {
	rangeDate := c.QueryParam("range")

	if rangeDate != "daily" && rangeDate != "weekly" && rangeDate != "monthly" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid range")
	}
	

	transactions, err := database.GetTransationsRangeDate(rangeDate)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   transactions,
	})
}