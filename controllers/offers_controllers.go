package controllers

import (
	"net/http"
	"vocaportal/services/offers"

	"github.com/labstack/echo/v4"
)

// Academic offers controller
type offersController struct{}

// Academic offers controller variable
var Offers = offersController{}

// Controller: /ofertas
func (oc offersController) FetchAll(c echo.Context) error {
	paramValues, httperr := extractParams(c, "programa", "universidad", "ciudad", "pais")

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	offers, httperr := offers.FetchAll(paramValues)

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, offers)
}
