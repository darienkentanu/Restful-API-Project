package main

import (
	"altastore/config"
	"altastore/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}