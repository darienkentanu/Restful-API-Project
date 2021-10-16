package routes

import (
	"altastore/controllers"
	"altastore/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", controllers.LoginUsersController)
	e.POST("/register", controllers.CreateUsersController)

	// Admin
	e.GET("/users", controllers.GetAllUsersController, middlewares.IsLoggedIn, middlewares.IsAdmin)
	
	// Current login user
	e.PUT("/users/:id", controllers.UpdateUserController, middlewares.IsLoggedIn)
	
	return e
}