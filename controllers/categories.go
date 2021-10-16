package controllers

import (
	"altastore/lib/database"
	"altastore/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCategories(c echo.Context) error {
	categories, err := database.GetCategories()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var res models.Category_response
	var resSlc []models.Category_response
	for _, v := range categories {
		res.ID = v.ID
		res.Name = v.Name
		resSlc = append(resSlc, res)
	}
	return c.JSON(http.StatusOK, resSlc)
}

func AddCategories(c echo.Context) error {
	var category models.Category
	c.Bind(&category)

	_, err := database.InsertCategories(category)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data": M{
			"id":   category.ID,
			"name": category.Name,
		},
	})
}

func DeleteCategories(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = database.DeleteCategoriesById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, M{
		"message": "category succesfully deleted",
	})
}
