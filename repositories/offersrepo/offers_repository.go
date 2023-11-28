package offersrepo

import (
	"vocaportal/models"
	"vocaportal/repositories"
)

// SQL Queries
const (
	fetchAllSQL = `
	SELECT
	Offer.ID,
	Offer.NAME,
	Programm.ID,
	Programm.NAME,
	University.ID,
	University.NAME,
	University.URL,
	City.ID,
	City.NAME,
	Country.ID,
	Country.NAME
	FROM
	academic_offers as Offer
	JOIN academic_programmes as Programm ON Offer.AP_ID = Programm.ID
	JOIN universities as University ON Offer.UNIVERSITY_ID = University.ID
	JOIN cities as City ON University.CITY_ID = City.ID
	JOIN countries as Country ON City.COUNTRY_ID = Country.ID
	`

	searchByNameSQL = `
	SELECT
	Offer.ID,
	Offer.NAME,
	Programm.ID,
	Programm.NAME,
	University.ID,
	University.NAME,
	University.URL,
	City.ID,
	City.NAME,
	Country.ID,
	Country.NAME
	FROM
	academic_offers as Offer
	JOIN academic_programmes as Programm ON Offer.AP_ID = Programm.ID
	JOIN universities as University ON Offer.UNIVERSITY_ID = University.ID
	JOIN cities as City ON University.CITY_ID = City.ID
	JOIN countries as Country ON City.COUNTRY_ID = Country.ID
	WHERE Programm.NAME LIKE $1 OR Offer.NAME LIKE $1
	`
)

// Stmt Creator
var stmtCreator = repositories.NewStmtCreator("academic offers")

// Stmts
var (
	fetchAllStmt     = stmtCreator.NewStmt(fetchAllSQL)
	searchByNameStmt = stmtCreator.NewStmt(searchByNameSQL)
)

// Fetch All
func FetchAll() ([]*models.Offer, error) {
	r, err := repositories.DoSimpleQuery(fetchAllStmt)

	if err != nil {
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateOffer)
}

// Search by name
func SearchByName(name string) ([]*models.Offer, error) {
	r, err := repositories.DoSimpleQuery(searchByNameStmt, name)

	if err != nil {
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateOffer)
}
