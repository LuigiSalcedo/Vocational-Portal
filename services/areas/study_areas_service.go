package areas

import (
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/areasrepo"
)

// Fetch all areas
func FetchAll() ([]*models.Area, *core.HttpError) {
	areas, err := areasrepo.FetchAll()

	if err != nil {
		return nil, core.InternalError
	}

	return areas, nil
}
