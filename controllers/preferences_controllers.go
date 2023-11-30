package controllers

import (
	"net/http"
	"vocaportal/core"
	"vocaportal/services/preferences"

	"github.com/labstack/echo/v4"
)

// Preferences controller struct
type preferencesController struct{}

// Prefences controller variable
var Preferences = preferencesController{}

// Fetch All
func (pc preferencesController) FetchAll(c echo.Context) error {
	p, httperr := preferences.FetchAll()

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, p)
}

// Search By Name
func (pc preferencesController) SearchByName(c echo.Context) error {
	name := c.Param("name")

	if len(name) == 0 {
		return c.JSON(core.PathError.Code, core.PathError)
	}

	p, httperr := preferences.SearchByName(name)

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, p)
}
