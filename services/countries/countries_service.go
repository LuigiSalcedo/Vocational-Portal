package countries

import (
	"strings"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/countriesrepo"
	"vocaportal/services"
)

// Service func to search countries by name
func SearchByName(name string) ([]*models.Country, *core.HttpError) {
	name = strings.ToUpper("%" + name + "%")

	countries, err := countriesrepo.SearchByName(name)

	if err != nil {
		return nil, core.InternalError
	}

	services.SortWithPrefix(countries, strings.Trim(name, "%"))

	return countries, nil
}

// Service func to fetch a country
func FetchCountry(id int64) (*models.Country, *core.HttpError) {
	c, err := countriesrepo.FetchCountry(id)

	if err != nil {
		return nil, core.InternalError
	}

	if c == nil {
		return nil, core.NotFoundError
	}

	return c, nil
}

// Services to fetch all countries
func FetchAll() ([]*models.Country, *core.HttpError) {
	c, err := countriesrepo.FetchAll()

	if err != nil {
		return nil, core.InternalError
	}

	return c, nil
}
