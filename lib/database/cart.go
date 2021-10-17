package database

import (
	"altastore/config"
	"altastore/models"
)

func CreateCart(cart models.Cart) error {
	if err := config.InitDB().Save(&cart).Error; err != nil {
		return err
	}

	return nil
}

func GetUserIdInCart(cartID int) (models.Cart, error) {
	var cart models.Cart
	if err := config.InitDB().Find(&cart, "id = ?", cartID).Error; err != nil {
		return cart, err
	}

	return cart, nil
}