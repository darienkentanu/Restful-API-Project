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

	e.POST("/carts", controllers.AddCartItemController, middlewares.IsLoggedIn)					// by user id
	e.GET("/carts", controllers.GetAllCartItem, middlewares.IsLoggedIn)							// by user id
	e.PUT("/cartitems/:id", controllers.UpdateCartItemController, middlewares.IsLoggedIn)		// by cart item id
	e.DELETE("/cartitems/:id", controllers.DeleteCartItemController, middlewares.IsLoggedIn)	// by cart item id

	e.POST("/checkout", controllers.Checkout, middlewares.IsLoggedIn)

	e.POST("/payments/:id", controllers.RequestBilling, middlewares.IsLoggedIn)
	e.GET("/payments/:id", controllers.GetTransactionStatus, middlewares.IsLoggedIn)

	e.GET("/transactions", controllers.GetAllTransactionsController, middlewares.IsLoggedIn)
	e.GET("/transactionsreport", controllers.GetTransactionsWithRangeDate, middlewares.IsLoggedIn, middlewares.IsAdmin)	// daily, weekly, monthly


	return e
}
