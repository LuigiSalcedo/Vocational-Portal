package controllers

import (
	"strconv"
	"vocaportal/core"

	"github.com/labstack/echo/v4"
)

// Auxiliar func to extract params from request
func extractParams(c echo.Context, params ...string) ([]int64, *core.HttpError) {
	values := make([]int64, 0, len(params))

	for _, param := range params {
		valueParam := c.QueryParam(param)
		value, err := strconv.Atoi(valueParam)

		if err != nil && len(valueParam) != 0 {
			return nil, core.NewQueryError(param)
		} else if len(valueParam) == 0 {
			value = -1
		}

		values = append(values, int64(value))
	}

	return values, nil
}
