package controllers

import (
	"net/http"
	"strconv"
	"vocaportal/core"
	"vocaportal/services/universities"

	"github.com/labstack/echo/v4"
)

// Universities struct controller
type universitiesController struct{}

// Universities controller variable
var Universities = universitiesController{}

// Controller: /universidades/nombre/:name
func (uc universitiesController) SearchByName(c echo.Context) error {
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

// Controller: /universidades/id/:id[0-9]+
func (uc universitiesController) FetchUniversity(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(core.PathError.Code, core.PathError)
	}

	u, httperr := universities.FetchUniversity(int64(id))

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, u)
}

// Controller: /universidades
func (uc universitiesController) FetchAll(c echo.Context) error {
	ciudadParam := c.QueryParam("ciudad")
	paisParam := c.QueryParam("pais")

	u, httperr := universities.FetchAll(ciudadParam, paisParam)

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, u)
}
