package controllers

import (
	"net/http"
	"vocaportal/services/programms"

	"github.com/labstack/echo/v4"
)

// Academic programms controller struct
type programmsController struct{}

// Academic programms varible
var Programms = programmsController{}

// Search by name
func (pc programmsController) SearchByName(c echo.Context) error {
	name := c.Param("name")

	programms, httperr := programms.SearchByName(name)

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, programms)
}
