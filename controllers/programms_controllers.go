package controllers

import (
	"net/http"
	"strconv"
	"vocaportal/core"
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

// Fetch by id
func (pc programmsController) FetchProgramm(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(core.PathError.Code, core.PathError)
	}

	p, httperr := programms.FetchProgramm(int64(id))

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, p)
}
