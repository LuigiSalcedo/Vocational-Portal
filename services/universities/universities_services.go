package universities

import (
	"strings"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/universitiesrepo"
)

// Service to search universities by name
func SearchByName(name string) ([]models.Univerisity, *core.HttpError) {
	name = strings.ToUpper(name)

	unis, err := universitiesrepo.SearchByName("%" + name + "%")

	if err != nil {
		return nil, core.InternalError
	}

	return unis, nil
}
