package main

import (
	"log"
	"vocaportal/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "static")

	// Endpoints universidades
	e.GET("/universidades/nombre/:name", controllers.Universities.SearchByName)

	log.Fatal(e.Start(":8088")) // Initializing server
}
