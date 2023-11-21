package controllers

import (
	"encoding/json"
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

// Controller: /programas/nombre/:nombre
func (pc programmsController) SearchByName(c echo.Context) error {
	name := c.Param("name")

	programms, httperr := programms.SearchByName(name)

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, programms)
}

// Controller: /programas/id/:id
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

// Controller: /programas
func (pc programmsController) FetchAll(c echo.Context) error {
	p, httperr := programms.FetchAll()

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, p)
}

// Controller: /programas/areas
func (pc programmsController) SearchByAreaRelation(c echo.Context) error {
	data, httperr := extractParams(c, "precision")

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	precision := data[0]

	err := json.NewDecoder(c.Request().Body).Decode(&data)

	if err != nil {
		return c.JSON(core.JsonError.Code, core.JsonError)
	}

	programms, httperr := programms.SearchByAreaRelation(data, precision)

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, programms)
}
