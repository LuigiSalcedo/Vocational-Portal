package offers

import (
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/offersrepo"
)

// Service to fetch all offer register
func FetchAll() ([]*models.Offer, *core.HttpError) {
	offers, err := offersrepo.FetchAll()

	if err != nil {
		return nil, core.InternalError
	}

	return offers, nil
}
