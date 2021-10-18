package database

import (
	"altastore/config"
	"altastore/models"
)

func UpdateCheckoutIDinCartItem(cartitemsID int, checkout_id int) (int, error) {
	if err := config.InitDB().Where("id = ?", cartitemsID).Set("checkout_id", checkout_id).Error; err != nil {
		return 0, err
	}
	return checkout_id, nil
}

func AddCheckoutID() (int, error) {
	var checkout models.Checkout
	if err := config.InitDB().Save(&checkout).Error; err != nil {
		return 0, err
	}
	return checkout.ID, nil
}
