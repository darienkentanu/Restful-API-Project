package controllers

import (
	"altastore/lib/database"
	"altastore/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllCartItem (c echo.Context) error {
	// Cart ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1{
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	userID := GetUserIdController(id)
	if userID != middlewares.CurrentLoginUser(c) {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}
	
	cartItems, err := database.GetAllCartItem(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   cartItems,
	})
}

func GetUserIdController(cartID int) int {
	cart, err := database.GetUserIdInCart(cartID)
	if err != nil {
		return -1
	}

	return cart.UserID
}