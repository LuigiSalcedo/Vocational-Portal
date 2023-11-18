package controllers

import (
	"net/http"
	"strconv"
	"vocaportal/core"
	"vocaportal/services/countries"

	"github.com/labstack/echo/v4"
)

// Countries struct controllers
type countriesController struct{}

// Countries controllers variable
var Countries = countriesController{}

// Controller: /paises/nombre/:name
func (cc countriesController) SearchByName(c echo.Context) error {
	name := c.Param("name")

	if len(name) == 0 {
		return c.JSON(core.BadRequest.Code, core.BadRequest)
	}

	countries, err := countries.SearchByName(name)

	if err != nil {
		return c.JSON(core.InternalError.Code, core.InternalError)
	}

	return c.JSON(http.StatusOK, countries)
}

// Controller: /paises/id/:id
func (cc countriesController) FetchCountry(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(core.PathError.Code, core.PathError)
	}

	country, httperr := countries.FetchCountry(int64(id))

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, country)
}

// Controller: /paises
func (cc countriesController) FetchAll(c echo.Context) error {
	countries, httperr := countries.FetchAll()

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, countries)
}
