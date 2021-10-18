package database

import (
	"altastore/config"
	"altastore/models"
)

func CreateCartItem(CartID int, addItem models.AddCartItem) (models.CartItem, error) {
	var cartItem models.CartItem

	cartItem.CartID = CartID
	cartItem.ProductID = addItem.ProductID
	cartItem.Quantity = addItem.Quantity

	if err := config.InitDB().Select("cart_id", "product_id", "quantity").Create(&cartItem).Error; err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func GetProductInCartItem(cartID, productID int) int {
	var cartItem models.CartItem
	row := config.InitDB().Where("cart_id = ? and product_id = ?", cartID, productID).Find(&cartItem).RowsAffected
	return int(row)
}

func UpdateStockCartItem(cartID int, addItem models.AddCartItem) (models.CartItem, error) {
	var cartItem models.CartItem

	if err := config.InitDB().Where("cart_id = ? and product_id = ?", cartID, addItem.ProductID).First(&cartItem).Error; err != nil {
		return cartItem, err
	}

	updatedQuantity := addItem.Quantity + cartItem.Quantity

	if err := config.InitDB().Model(&cartItem).Update("quantity", updatedQuantity).Error; err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func GetAllCartItem(cartId int) ([]models.CartItem, error) {
	var cartItem []models.CartItem

	if err := config.InitDB().Find(&cartItem, "cart_id = ? and checkout_id IS NULL", cartId).Error; err != nil {
		return nil, err
	}

	return cartItem, nil
}

func UpdateCartItem(cartItemID int, updateItem models.UpdateCartItem) (models.CartItem, error) {
	var cartItem models.CartItem

	if err := config.InitDB().Where("id = ?", cartItemID).First(&cartItem).Error; err != nil {
		return cartItem, err
	}

	cartItem.Quantity = updateItem.Quantity

	if err := config.InitDB().Model(&cartItem).Update("quantity", cartItem.Quantity).Error; err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func DeleteCartItem(id int) error {
	var cartItem models.CartItem

	if err := config.InitDB().Where("id = ?", id).Delete(&cartItem).Error; err != nil {
		return err
	}

	return nil
}

func GetCartIdInCartItem(cartItemID int) (models.CartItem, error) {
	var cartItem models.CartItem
	if err := config.InitDB().Find(&cartItem, "id = ?", cartItemID).Error; err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func GetProductQuantityInCartItem(cartID, productID int) (models.CartItem, error) {
	var cartItem models.CartItem
	if err := config.InitDB().Where("cart_id = ? and product_id = ? and checkout_id IS NULL", cartID, productID).First(&cartItem).Error; err != nil {
		return cartItem, err
	}

	return cartItem, nil
}
