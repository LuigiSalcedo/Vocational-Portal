package offers

import (
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories/offersrepo"
	"vocaportal/services"
)

// Service to fetch all offer register
func FetchAll(paramValues []int64) ([]*models.Offer, *core.HttpError) {
	offers, err := offersrepo.FetchAll()

	if err != nil {
		return nil, core.InternalError
	}

	filters := make([]services.Filter[models.Offer], 0, len(paramValues))

	// Filtro por id de programa
	filters = services.CreateAndAddFilter(filters, func(o *models.Offer) any { return o.Programm.Id }, paramValues[0])

	// Filtro por id de universidad
	filters = services.CreateAndAddFilter(filters, func(o *models.Offer) any { return o.University.Id }, paramValues[1])

	// Filtro por id de ciudad
	filters = services.CreateAndAddFilter(filters, func(o *models.Offer) any { return o.University.City.Id }, paramValues[2])

	// Filtro por id de pais
	filters = services.CreateAndAddFilter(filters, func(o *models.Offer) any { return o.University.City.Country.Id }, paramValues[3])

	// Filtrar información
	return services.FilterData(offers, filters...), nil
}
