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
	e.GET("/universidades/id/:id", controllers.Universities.FetchUniversity)
	e.GET("/universidades", controllers.Universities.FetchAll)

	// Endpoints de paises
	e.GET("/paises/nombre/:name", controllers.Countries.SearchByName)
	e.GET("/paises/id/:id", controllers.Countries.FetchCountry)

	// Endpoints de ciudades
	e.GET("/ciudades/nombre/:name", controllers.Cities.SearchByName)
	e.GET("/ciudades/id/:id", controllers.Cities.FetchCity)

	log.Fatal(e.Start(":8088")) // Initializing server
}
