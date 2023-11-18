package cities

import (
	"strings"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/citiesrepo"
)

// Services to get cities by name
func SearchByName(name string) ([]models.City, *core.HttpError) {
	name = strings.ToUpper("%" + name + "%")

	cities, err := citiesrepo.SearchByName(name)

	if err != nil {
		return nil, core.InternalError
	}

	return cities, nil
}
