package controllers

import (
	"altastore/lib/database"
	"altastore/middlewares"
	"altastore/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddCartItemController(c echo.Context) error {
	var addItem models.AddCartItem

	// Cart ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1{
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	userID := GetUserIdController(id)
	if userID != middlewares.CurrentLoginUser(c) {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	if err := c.Bind(&addItem); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	if addItem.Quantity <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid quantity")
	}

	row := database.GetProductID(addItem.ProductID)
	if row == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid product id")
	}

	var cartItem models.CartItem
	row = database.GetProductInCartItem(id, addItem.ProductID)
	if row == 0 {
		cartItem, err = database.CreateCartItem(id, addItem)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
	} else {
		cartItem, err = database.UpdateStockCartItem(id, addItem)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   cartItem,
	})
}

func UpdateCartItemController(c echo.Context) error {
	var updateItem models.UpdateCartItem

	// Cart item ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1{
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	cartID := GetCartIdController(id)
	userID := GetUserIdController(cartID)
	if userID != middlewares.CurrentLoginUser(c) {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	if err := c.Bind(&updateItem); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	if updateItem.Quantity < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid quantity")
	}
	
	var cartItem models.CartItem
	if updateItem.Quantity == 0 {
		err = database.DeleteCartItem(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}

		return c.JSON(http.StatusOK, M{
			"message": "cart item succesfully deleted",
		})
	} else {
		cartItem, err = database.UpdateCartItem(id, updateItem)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   cartItem,
	})
}

func DeleteCartItemController(c echo.Context) error {
	// Cart item ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1{
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	cartID := GetCartIdController(id)
	userID := GetUserIdController(cartID)
	if userID != middlewares.CurrentLoginUser(c) {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	err = database.DeleteCartItem(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, M{
		"message": "cart item succesfully deleted",
	})
}

func GetCartIdController(cartItemID int) int {
	cartItem, err := database.GetCartIdInCartItem(cartItemID)
	if err != nil {
		return -1
	}

	return cartItem.CartID
}