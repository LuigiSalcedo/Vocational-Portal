package cities

import (
	"strings"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/citiesrepo"
	"vocaportal/services"
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

// Services to get all cities
func FetchAll(countryId int64) ([]*models.City, *core.HttpError) {
	cities, err := citiesrepo.FetchAll()

	if err != nil {
		return nil, core.InternalError
	}

	if countryId != -1 {
		filter := services.NewSimpleFilter(func(c *models.City) any { return c.Country.Id }, countryId)
		citiesFiltered := services.FilterData(cities, filter)
		cities = citiesFiltered
	}

	return cities, nil
}
