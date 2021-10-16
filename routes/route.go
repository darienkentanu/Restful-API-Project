package routes

import (
	"altastore/controllers"
	"altastore/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/register", controllers.CreateUsersController)
	e.POST("/login", controllers.LoginUsersController)

	// Admin
	e.GET("/users", controllers.GetAllUsersController, middlewares.IsLoggedIn, middlewares.IsAdmin)

	// Current login user
	e.PUT("/users/:id", controllers.UpdateUserController, middlewares.IsLoggedIn)

	e.GET("/categories", controllers.GetCategories, middlewares.IsLoggedIn)
	e.POST("/categories", controllers.AddCategories, middlewares.IsLoggedIn, middlewares.IsAdmin)
	e.DELETE("/categories/:id", controllers.DeleteCategories, middlewares.IsLoggedIn, middlewares.IsAdmin)
	return e
}
