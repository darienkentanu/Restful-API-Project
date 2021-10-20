package controllers

import (
	"altastore/lib/database"
	"altastore/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProductsController(c echo.Context) error {
	var product models.Product

	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	err := database.GetCategoryId(product.CategoryID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category id")
	}

	product, err = database.CreateProduct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusCreated, M{
		"status": "success",
		"data":   product,
	})
}

func GetAllProductsController(c echo.Context) error {
	products, err := database.GetAllProducts()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   products,
	})
}

func GetProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	product, err := database.GetProduct(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   product,
	})
}

func UpdateProductController(c echo.Context) error {
	var newProduct models.Product

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	if err := c.Bind(&newProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	product, err := database.UpdateProduct(id, newProduct)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   product,
	})
}

func DeleteProductController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	err = database.DeleteProduct(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, M{
		"message": "product succesfully deleted",
	})
}
