package controllers

import (
	"altastore/lib/database"
	"altastore/middlewares"
	"altastore/models"
	"altastore/payment"
	"net/http"

	"github.com/labstack/echo/v4"
	gubrak "github.com/novalagung/gubrak/v2"
)

func Checkout(c echo.Context) error {
	var checkoutItems models.CheckoutItems_Response
	var inputJsonCheckout models.CheckoutItems_Input
	c.Bind(&inputJsonCheckout)

	userID := middlewares.CurrentLoginUser(c)
	cartID := CartIdInCart(userID)
	// checkoutItems.CartID = cartID
	cartItems, err := database.GetAllCartItem(cartID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	// hasil dari list product yang dipilih berdasarkan idnya
	cartItemsSelected := inputJsonCheckout.ProductID
	var productIDSelected []int //  berupa product id yg dipilih
	for _, k := range cartItems {
		itemInCartItem := k
		for _, v := range cartItemsSelected {
			if itemInCartItem.ProductID == v {
				productIDSelected = append(productIDSelected, v)
			}
		}
	}

	// ambil harga product dari product && nama product by id

	var harga = make(map[int]int)
	var productName = make(map[int]string)
	for _, v := range productIDSelected {
		product, err := database.GetProduct(v)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		harga[product.ID] = product.Price
		productName[product.ID] = product.Name
	}

	// kalikan harga dengan quantity
	var amount int
	for i, v := range harga {
		product, err := database.GetProductQuantityInCartItem(cartID, i)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		amount += product.Quantity * v
	}

	// buat orderID
	orderId := gubrak.RandomInt(10000, 99999)

	// buat checkoutID
	checkoutItems.ID, err = database.AddCheckoutID()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	var cartItemRes models.CartItem_Response
	for _, v := range productIDSelected {
		cartItemRes.ProductID = v
		cartItemRes.ProductName = productName[v]
		product, err := database.GetProductQuantityInCartItem(cartID, v)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		cartItemRes.Quantity = product.Quantity
		cartItemRes.ProductPrice = harga[product.ProductID]
		checkoutItems.Product = append(checkoutItems.Product, cartItemRes)
	}

	// insert checkoutid to cartitems
	database.UpdateCheckoutIDinCartItem(cartID, checkoutItems.ID)

	// request payment
	redirectURL, err := payment.RequestBilling(orderId, amount)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, "error request ke midtrans")
	}
	// add to transactions
	err = database.AddTransaction(inputJsonCheckout.Courier, orderId, amount, userID, checkoutItems.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, M{
		"data":            checkoutItems,
		"orderid":         orderId,
		"total amount":    amount,
		"status":          "success",
		"link pembayaran": redirectURL,
	})
}
