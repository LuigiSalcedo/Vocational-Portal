package preferences

import (
	"strings"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/preferencesrepo"
	"vocaportal/services"
)

func FetchAll() ([]*models.Preference, *core.HttpError) {
	p, err := preferencesrepo.FetchAll()

	if err != nil {
		return nil, core.InternalError
	}

	return p, nil
}

func SearchByName(name string) ([]*models.Preference, *core.HttpError) {
	name = strings.ToUpper("%" + name + "%")
	p, err := preferencesrepo.SearchByName(name)

	if err != nil {
		return nil, core.InternalError
	}

	services.SortWithPrefix(p, name)

	return p, nil
}
