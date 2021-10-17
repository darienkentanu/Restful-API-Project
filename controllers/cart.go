package controllers

import (
	"altastore/lib/database"
	"altastore/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllCartItem (c echo.Context) error {
	// User ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1{
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	if id != middlewares.CurrentLoginUser(c) {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}
	
	cartID := CartIdInCart(id)
	cartItems, err := database.GetAllCartItem(cartID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   cartItems,
	})
}

func UserIdInCart(cartID int) int {
	cart, err := database.GetUserIdInCart(cartID)
	if err != nil {
		return -1
	}

	return cart.UserID
}

func CartIdInCart(userID int) int {
	cart, err := database.GetCartIdInCart(userID)
	if err != nil {
		return -1
	}

	return cart.ID
}