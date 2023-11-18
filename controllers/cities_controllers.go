package controllers

import (
	"net/http"
	"vocaportal/core"
	"vocaportal/services/cities"

	"github.com/labstack/echo/v4"
)

// Cities controller struct
type citiesController struct{}

// Cities controllers variable
var Cities = citiesController{}

func (cc citiesController) SearchByName(c echo.Context) error {
	name := c.Param("name")

	r, err := cities.SearchByName(name)

	if err != nil {
		return c.JSON(core.InternalError.Code, core.InternalError)
	}

	return c.JSON(http.StatusOK, r)
}
