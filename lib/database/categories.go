package database

import (
	"altastore/config"
	"altastore/models"
	"errors"
)

func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := config.InitDB().Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

func InsertCategories(categories models.Category) (interface{}, error) {
	if err := config.InitDB().Save(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func DeleteCategoriesById(id int) error {

	rows := config.InitDB().Delete(&models.Category{}, id).RowsAffected
	if rows == 0 {
		err := errors.New("categories to be deleted is not found")
		return err
	}
	return nil

}
