package main

import (
	"altastore/config"
	"altastore/middlewares"
	"altastore/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
