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
