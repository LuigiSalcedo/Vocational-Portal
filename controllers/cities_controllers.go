package controllers

import (
	"net/http"
	"strconv"
	"vocaportal/core"
	"vocaportal/services/cities"

	"github.com/labstack/echo/v4"
)

// Cities controller struct
type citiesController struct{}

// Cities controllers variable
var Cities = citiesController{}

// Controller: /ciudades/nombre/:name
func (cc citiesController) SearchByName(c echo.Context) error {
	name := c.Param("name")

	r, err := cities.SearchByName(name)

	if err != nil {
		return c.JSON(core.InternalError.Code, core.InternalError)
	}

	return c.JSON(http.StatusOK, r)
}

// Controller: /ciudades/id/:id
func (cc citiesController) FetchCity(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(core.PathError.Code, core.PathError)
	}

	city, httperr := cities.FetchCity(int64(id))

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, city)
}

// Controller: /ciudades
func (cc citiesController) FetchAll(c echo.Context) error {
	countryIdParam := c.QueryParam("pais")

	countryId, err := strconv.Atoi(countryIdParam)

	if err != nil && len(countryIdParam) != 0 {
		return c.JSON(core.QueryError.Code, core.NewQueryError("pais"))
	}

	if len(countryIdParam) == 0 {
		countryId = -1
	}

	cities, httperr := cities.FetchAll(int64(countryId))

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, cities)
}
