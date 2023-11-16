package main

import (
	"log"
	"vocaportal/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "static")

	// Ver documentación
	e.GET("/documentation", controllers.Documentation.APIDoc)

	// Endpoints universidades
	e.GET("/universidades/nombre/:name", controllers.Universities.SearchByName)
	e.GET("/universidades/id/:id", controllers.Universities.SerachById)

	// Endpoints de paises
	e.GET("/paises/nombre/:name", controllers.Countries.SearchByName)

	log.Fatal(e.Start(":8088")) // Initializing server
}
