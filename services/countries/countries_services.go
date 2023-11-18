package countries

import (
	"strings"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/countriesrepo"
)

// Service func to search countries by name
func SearchByName(name string) ([]*models.Country, *core.HttpError) {
	name = strings.ToUpper("%" + name + "%")

	countries, err := countriesrepo.SearchByName(name)

	if err != nil {
		return nil, core.InternalError
	}

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
