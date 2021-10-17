package routes

import (
	"altastore/controllers"
	"altastore/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	middlewares.LogMiddlewares(e)
	e.POST("/register", controllers.CreateUsersController)
	e.POST("/login", controllers.LoginUsersController)

	e.GET("/users", controllers.GetAllUsersController, middlewares.IsLoggedIn, middlewares.IsAdmin)
	e.PUT("/users/:id", controllers.UpdateUserController, middlewares.IsLoggedIn)

	e.GET("/categories", controllers.GetCategories, middlewares.IsLoggedIn)
	e.POST("/categories", controllers.AddCategories, middlewares.IsLoggedIn, middlewares.IsAdmin)
	e.DELETE("/categories/:id", controllers.DeleteCategories, middlewares.IsLoggedIn, middlewares.IsAdmin)

	e.GET("/products", controllers.GetAllProductsController)
	e.GET("/products/:id", controllers.GetProductController)
	e.POST("/products", controllers.CreateProductsController, middlewares.IsLoggedIn, middlewares.IsAdmin)
	e.PUT("/products/:id", controllers.UpdateProductController, middlewares.IsLoggedIn, middlewares.IsAdmin)
	e.DELETE("/products/:id", controllers.DeleteProductController, middlewares.IsLoggedIn, middlewares.IsAdmin)

	return e
}
