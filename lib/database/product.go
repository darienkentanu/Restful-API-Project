package database

import (
	"altastore/config"
	"altastore/models"
)

func CreateProduct(product models.Product) (models.Product, error) {
	if err := config.InitDB().Save(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	if err := config.InitDB().Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func GetProduct(id int) (models.Product, error) {
	var product models.Product

	if err := config.InitDB().Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func UpdateProduct(id int, newProduct models.Product) (models.Product, error) {
	var product models.Product

	if err := config.InitDB().First(&product, id).Error; err != nil {
		return product, err
	}

	product.Name 		= newProduct.Name
	product.CategoryID 	= newProduct.CategoryID
	product.Description = newProduct.Description
	product.Quantity 	= newProduct.Quantity
	product.Price 		= newProduct.Price

	if err := config.InitDB().Save(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func DeleteProduct(id int) error {
	var product models.Product

	if err := config.InitDB().Where("id = ?", id).Delete(&product).Error; err != nil {
		return err
	}

	return nil
}