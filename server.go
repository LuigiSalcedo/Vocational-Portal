package main

import (
	"crypto/subtle"
	"log"
	"os"
	"vocaportal/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Configuración de CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{"GET", "POST"},
	}))

	e.Static("/", "static")

	// Ver documentación
	e.GET("/documentation", controllers.Documentation.APIDoc)

	g := e.Group("")

	// Basic Auth para acceder a la vista de admin
	g.Use(middleware.BasicAuth(func(user, password string, ctx echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(user), []byte(os.Getenv("vocaportal_admin"))) != 1 {
			return false, nil
		}

		if subtle.ConstantTimeCompare([]byte(password), []byte(os.Getenv("vocaportal_password"))) != 1 {
			return false, nil
		}

		return true, nil
	}))

	// Cargar vista de admins
	g.GET("/admin", controllers.Admin.LoadHTML)

	// Almacenar un administrador
	e.POST("/admins/guardar", controllers.Admin.SaveAdmin)

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
	e.POST("/programas/areas", controllers.Programms.SearchByAreaRelation)
	e.POST("/programas/preferencias", controllers.Programms.SearchByPreferences)

	// Endpoint de ofertas
	e.GET("/ofertas", controllers.Offers.FetchAll)
	e.GET("/ofertas/nombre/:name", controllers.Offers.SearchByName)

	// Endpoint de areas de estudio
	e.GET("/areas", controllers.Areas.FetchAll)
	e.GET("/areas/nombre/:name", controllers.Areas.SearchByName)

	// Endpoints de preferencias
	e.GET("/preferencias", controllers.Preferences.FetchAll)
	e.GET("/preferencias/nombre/:name", controllers.Preferences.SearchByName)

	log.Fatal(e.Start(":8088")) // Initializing server
}
