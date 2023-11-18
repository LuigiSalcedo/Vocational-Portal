package programms

import (
	"strings"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/programmsrepo"
)

// Service to search academics programms by name
func SearchByName(name string) ([]*models.Programm, *core.HttpError) {
	name = strings.ToUpper("%" + name + "%")
	programms, err := programmsrepo.SearchByName(name)

	if err != nil {
		return nil, core.InternalError
	}

	return programms, nil
}
