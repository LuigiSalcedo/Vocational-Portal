package countries

import (
	"strings"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/countriesrepo"
)

// Services func to search countries by name
func SearchByName(name string) ([]models.Country, *core.HttpError) {
	name = strings.ToUpper("%" + name + "%")

	countries, err := countriesrepo.SearchByName(name)

	if err != nil {
		return nil, core.InternalError
	}

	return countries, nil
}
