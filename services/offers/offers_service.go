package offers

import (
	"strconv"
	"strings"
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

	if paramValues[4] != -1 {
		filters = append(filters, func(o *models.Offer) bool {
			if len(o.Price) <= 1 {
				return false
			}

			priceStr := strings.Trim(o.Price, "$")
			priceStr = strings.Replace(priceStr, ".", "", -1)

			price, err := strconv.Atoi(priceStr)

			if err != nil {
				core.LogErr(err)
				return false
			}

			return int64(price) <= paramValues[4]
		})
	}

	// Filtrar información
	r := services.FilterData(offers, filters...)

	if paramValues[4] != 1 {
		services.SortSlice(r, func(o1, o2 *models.Offer) int {
			return strings.Compare(o1.Price, o2.Price)
		})
	} else {
		services.SortSlice(r, func(o1, o2 *models.Offer) int {
			return strings.Compare(o1.Name, o2.Name)
		})
	}

	for _, offer := range offers {
		if offer.Price == "$0" {
			offer.Price = "Universidad pública"
		}

		if offer.Price == "$" {
			offer.Price = "$-"
		}
	}

	return r, nil
}

// Service to search offers by name
func SearchByName(name string) ([]*models.Offer, *core.HttpError) {
	name = strings.ToUpper(name)

	offers, err := offersrepo.SearchByName("%" + name + "%")

	if err != nil {
		return nil, core.InternalError
	}

	services.SortWithPrefix(offers, name)

	for _, offer := range offers {
		if offer.Price == "$0" {
			offer.Price = "Universidad pública"
		}

		if offer.Price == "$" {
			offer.Price = "$-"
		}
	}

	return offers, nil
}
