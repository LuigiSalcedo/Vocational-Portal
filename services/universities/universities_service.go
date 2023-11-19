package universities

import (
	"strconv"
	"strings"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/universitiesrepo"
	"vocaportal/services"
)

// Service to search universities by name
func SearchByName(name string) ([]*models.University, *core.HttpError) {
	name = strings.ToUpper(name)

	unis, err := universitiesrepo.SearchByName("%" + name + "%")

	if err != nil {
		return nil, core.InternalError
	}

	services.SortWithPrefix(unis, name)

	return unis, nil
}

// Service to fetch a university by id
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

// Service to fetch every universities register
func FetchAll(ciudadParam, paisParam string) ([]*models.University, *core.HttpError) {
	u, err := universitiesrepo.FetchAll()

	if err != nil {
		return nil, core.InternalError
	}

	filters := make([]services.Filter[models.University], 0, 2)

	if len(ciudadParam) != 0 {
		ciudad, err := strconv.Atoi(ciudadParam)

		filters = append(filters, services.NewSimpleFilter(
			func(u *models.University) any { return u.City.Id }, int64(ciudad)),
		)

		if err != nil {
			return nil, core.NewQueryError("ciudad")
		}
	}

	if len(paisParam) != 0 {
		pais, err := strconv.Atoi(paisParam)

		filters = append(filters, services.NewSimpleFilter(
			func(u *models.University) any { return u.City.Country.Id }, int64(pais)),
		)

		if err != nil {
			return nil, core.NewQueryError("pais")
		}
	}

	return services.FilterData(u, filters...), nil
}
