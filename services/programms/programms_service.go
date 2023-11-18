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

// Service to fecth an academic programm by id
func FetchProgramm(id int64) (*models.Programm, *core.HttpError) {
	p, err := programmsrepo.FetchProgramm(id)

	if err != nil {
		return nil, core.InternalError
	}

	if p == nil {
		return nil, core.NotFoundError
	}

	return p, nil
}

// Service to fetch every academic programm register
func FetchAll() ([]*models.Programm, *core.HttpError) {
	p, err := programmsrepo.FetchAll()

	if err != nil {
		return nil, core.InternalError
	}

	return p, nil
}
