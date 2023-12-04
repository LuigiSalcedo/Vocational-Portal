package programms

import (
	"strings"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/areasrepo"
	"vocaportal/repositories/programmsrepo"
	"vocaportal/services"
)

// Service to search academics programms by name
func SearchByName(name string) ([]*models.Programm, *core.HttpError) {
	name = strings.ToUpper("%" + name + "%")
	programms, err := programmsrepo.SearchByName(name)

	if err != nil {
		return nil, core.InternalError
	}

	services.SortWithPrefix(programms, strings.Trim(name, "%"))

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

	services.SortSlice(p, func(p1, p2 *models.Programm) int {
		return strings.Compare(p1.Name, p2.Name)
	})

	return p, nil
}

// Services to get programms by area relation
func SearchByAreaRelation(areasId []int64, precision int64) ([]*models.Programm, *core.HttpError) {
	pwars, err := programmsrepo.SearchByAreaRelation(areasId)

	if err != nil {
		return nil, core.InternalError
	}

	filters := make([]services.Filter[models.PWAR], 0, 1)

	if precision > 0 {
		filters = append(filters, func(p *models.PWAR) bool {
			return p.N >= float64(precision)
		})
	}

	pwars = services.FilterData(pwars, filters...)

	programms := make([]*models.Programm, 0, len(pwars))

	for _, pwar := range pwars {
		programms = append(programms, &pwar.Programm)
	}

	return programms, nil
}

// Service to get programms by prefernce
func SearchByPreferences(preferencesIds []int64) ([]*models.Programm, *core.HttpError) {
	areasIds, err := areasrepo.SearchByPreference(preferencesIds)

	if err != nil {
		return nil, core.InternalError
	}

	return SearchByAreaRelation(areasIds, 0)
}
