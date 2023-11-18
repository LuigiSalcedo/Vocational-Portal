package controllers

import (
	"bufio"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"vocaportal/core"

	"github.com/labstack/echo/v4"
)

type docController struct{}

type Endpoint struct {
	Description  string
	Path         string
	Method       string
	JsonRequest  string
	JsonResponse string
}

func ToEndpoint(str string) Endpoint {
	data := strings.Split(str, "|")

	e := Endpoint{
		Description:  data[0],
		Path:         data[1],
		Method:       data[2],
		JsonRequest:  data[3],
		JsonResponse: data[4],
	}

	return e
}

var Documentation = docController{}

func (dC docController) APIDoc(c echo.Context) error {
	t, err := template.ParseFiles("./static/docs/doc_template.html")

	if err != nil {
		core.LogErr(err)
		return c.String(http.StatusInternalServerError, "Error abriendo la documentación")
	}

	data := make([]Endpoint, 0, 20)

	file, err := os.Open("./static/docs/documentation.apidoc")

	if err != nil {
		core.LogErr(err)
		return c.String(http.StatusInternalServerError, "Error abriendo la documentación")
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("Ocurrió un error abriendo la documentación\n.Se cargo hasta:\n", data[len(data)-1])
		}
	}()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		str := string(sc.Bytes())
		data = append(data, ToEndpoint(str))
	}

	parsedFile := bytes.NewBuffer(nil)
	err = t.Execute(parsedFile, data)

	if err != nil {
		core.LogErr(err)
		return c.String(http.StatusInternalServerError, "Error abriendo la documentación")
	}

	return c.HTML(200, parsedFile.String())
}
