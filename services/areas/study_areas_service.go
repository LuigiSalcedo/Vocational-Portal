package areas

import (
	"strings"
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/areasrepo"
	"vocaportal/services"
)

// Fetch all areas
func FetchAll() ([]*models.Area, *core.HttpError) {
	areas, err := areasrepo.FetchAll()

	if err != nil {
		return nil, core.InternalError
	}

	services.SortSlice(areas, func(a1, a2 *models.Area) int {
		return strings.Compare(a1.Name, a2.Name)
	})

	return areas, nil
}

// Fetch areas by name
func SearchByName(name string) ([]*models.Area, *core.HttpError) {
	name = strings.ToUpper(name)
	areas, err := areasrepo.SearchByName("%" + name + "%")

	if err != nil {
		return nil, core.InternalError
	}

	services.SortWithPrefix(areas, name)

	return areas, nil
}
