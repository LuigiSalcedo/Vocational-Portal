package controllers

import (
	"net/http"
	"vocaportal/core"
	"vocaportal/services/universities"

	"github.com/labstack/echo/v4"
)

// Universities struct controller
type universitiesController struct{}

// Universities controller variable
var Universities = universitiesController{}

// Controllers /universidades/nombre/:name
func (u universitiesController) SearchByName(c echo.Context) error {
	name := c.Param("name")

	if len(name) == 0 {
		return c.JSON(core.PathError.Code, core.PathError)
	}

	unis, httperr := universities.SearchByName(name)

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, unis)
}
