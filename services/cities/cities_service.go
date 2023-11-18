package cities

import (
	"strings"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/citiesrepo"
)

// Service to get cities by name
func SearchByName(name string) ([]*models.City, *core.HttpError) {
	name = strings.ToUpper("%" + name + "%")

	cities, err := citiesrepo.SearchByName(name)

	if err != nil {
		return nil, core.InternalError
	}

	return cities, nil
}

// Service to get a city by id
func FetchCity(id int64) (*models.City, *core.HttpError) {
	city, err := citiesrepo.FetchCity(id)

	if err != nil {
		return nil, core.InternalError
	}

	if city == nil {
		return nil, core.NotFoundError
	}

	return city, nil
}
