package database

import (
	"altastore/config"
	"altastore/models"
)

func AddCheckoutID() (int, error) {
	var checkout models.Checkout
	if err := config.InitDB().Save(&checkout).Error; err != nil {
		return 0, err
	}
	return checkout.ID, nil
}

func UpdateCheckoutIdInCartItem(checkoutID, cartID, productID int) (models.CartItem, error) {
	var cartItem models.CartItem

	if err := config.InitDB().Where("cart_id = ? and product_id = ? and checkout_id IS NULL", cartID, productID).First(&cartItem).Error; err != nil {
		return cartItem, err
	}

	cartItem.CheckoutID = checkoutID

	if err := config.InitDB().Model(&cartItem).Update("checkout_id", cartItem.CheckoutID).Error; err != nil {
		return cartItem, err
	}

	return cartItem, nil
}