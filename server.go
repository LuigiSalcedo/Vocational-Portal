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
	e.GET("/paises", controllers.Countries.FetchAll)

	// Endpoints de ciudades
	e.GET("/ciudades/nombre/:name", controllers.Cities.SearchByName)
	e.GET("/ciudades/id/:id", controllers.Cities.FetchCity)
	e.GET("/ciudades", controllers.Cities.FetchAll)

	// Endpoints de programas académicos
	e.GET("/programas/nombre/:name", controllers.Programms.SearchByName)
	e.GET("/programas/id/:id", controllers.Programms.FetchProgramm)
	e.GET("/programas", controllers.Programms.FetchAll)

	// Endpoint de ofertas
	e.GET("/ofertas", controllers.Offers.FetchAll)
	e.GET("/ofertas/nombre/:name", controllers.Offers.SearchByName)

	// Endpoint de areas de estudio
	e.GET("/areas", controllers.Areas.FetchAll)

	log.Fatal(e.Start(":8088")) // Initializing server
}
