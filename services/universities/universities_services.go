package universities

import (
	"strings"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/universitiesrepo"
)

// Service to search universities by name
func SearchByName(name string) ([]models.University, *core.HttpError) {
	name = strings.ToUpper(name)

	unis, err := universitiesrepo.SearchByName("%" + name + "%")

	if err != nil {
		return nil, core.InternalError
	}

	return unis, nil
}

// Services to fetch a university by id
func FetchUniversity(id int64) (*models.University, *core.HttpError) {
	u, err := universitiesrepo.FetchUniversity(id)

	if err != nil {
		return nil, core.InternalError
	}

	if u == nil {
		return nil, core.NotFoundError
	}

	return u, nil
}
