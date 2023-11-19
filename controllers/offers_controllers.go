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
	offers, httperr := offers.FetchAll()

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, offers)
}
