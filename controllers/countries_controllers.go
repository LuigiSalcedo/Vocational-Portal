package controllers

import (
	"net/http"
	"vocaportal/core"
	"vocaportal/services/countries"

	"github.com/labstack/echo/v4"
)

// Countries controllers
type countriesController struct{}

var Countries = countriesController{}

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
