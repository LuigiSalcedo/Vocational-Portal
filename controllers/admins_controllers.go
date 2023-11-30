package controllers

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/services/admins"

	"github.com/labstack/echo/v4"
)

// Admins controller struct
type adminsController struct{}

// Admins controller variable
var Admin = adminsController{}

// Controller: /admins/guardar
func (ac adminsController) SaveAdmin(c echo.Context) error {
	admin := models.Admin{}

	err := json.NewDecoder(c.Request().Body).Decode(&admin)

	if err != nil {
		return c.JSON(core.JsonError.Code, core.JsonError)
	}

	httperr := admins.SaveAdmin(admin)

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.NoContent(http.StatusCreated)
}

// Controller: /admins/entrar
func (ac adminsController) Login(c echo.Context) error {
	admin := models.Admin{}

	err := json.NewDecoder(c.Request().Body).Decode(&admin)

	if err != nil {
		return c.JSON(core.JsonError.Code, core.JsonError)
	}

	token, httperr := admins.Login(admin)

	if httperr != nil {
		return c.JSON(httperr.Code, httperr)
	}

	return c.JSON(http.StatusOK, token)
}

// Controller: /admin
func (ac adminsController) LoadHTML(c echo.Context) error {
	t, err := template.ParseFiles("./static/admin.html")

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error loading admin.html")
	}

	fileData := bytes.NewBuffer(nil)

	t.Execute(fileData, nil)

	return c.HTML(http.StatusOK, fileData.String())
}
