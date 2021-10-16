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