package controllers

import (
	"altastore/lib/database"
	"altastore/middlewares"
	"altastore/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Checkout(c echo.Context) error {
	var checkoutItems models.CheckoutItems

	userID := middlewares.CurrentLoginUser(c)

	cartID := CartIdInCart(userID)
	checkoutItems.CartID = cartID
	cartItems, err := database.GetAllCartItem(cartID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	checkoutItems.CartItem = cartItems

	checkoutItems.ID, err = database.AddCheckoutID()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   checkoutItems,
	})
}
