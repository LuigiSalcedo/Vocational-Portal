package controllers

import (
	"net/http"
	"vocaportal/services/areas"

	"github.com/labstack/echo/v4"
)

// Study areas controller struc
type areasController struct{}

// Study areas controllers variable
var Areas = areasController{}

// Controller: /areas
func (ac areasController) FetchAll(c echo.Context) error {
	areas, httperr := areas.FetchAll()

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, areas)
}

// Controllers: /areas/nombre/:name
func (ac areasController) SearchByName(c echo.Context) error {
	name := c.Param("name")

	areas, httperr := areas.SearchByName(name)

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, areas)
}
