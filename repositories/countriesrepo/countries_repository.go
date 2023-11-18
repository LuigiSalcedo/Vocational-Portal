package countriesrepo

import (
	"vocaportal/models"
	"vocaportal/repositories"
)

const (
	searchByNameSQL = `
	SELECT
	ID, 
	NAME
	FROM countries
	WHERE NAME LIKE $1
	`

	fetchCountrySQL = `
	SELECT
	ID,
	NAME
	FROM countries
	WHERE ID = $1 LIMIT 1
	`

	fetchAllSQL = `
	SELECT
	ID,
	NAME
	FROM countries
	`
)

// Stmt creator
var stmtCreator = repositories.NewStmtCreator("countries")

// Stmts
var (
	searchByNameStmt = stmtCreator.NewStmt(searchByNameSQL)
	fetchCountryStmt = stmtCreator.NewStmt(fetchCountrySQL)
	fetchAllStmt     = stmtCreator.NewStmt(fetchAllSQL)
)

func SearchByName(name string) ([]*models.Country, error) {
	r, err := repositories.DoSimpleQuery(searchByNameStmt, name)

	if err != nil {
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateCountry)
}

func FetchCountry(id int64) (*models.Country, error) {
	r, err := repositories.DoSimpleQuery(fetchCountryStmt, id)

	if err != nil {
		return nil, err
	}

	result, err := repositories.Data(r, models.CreateCountry)

	if result == nil {
		return nil, err
	}

	return *result, err
}

func FetchAll() ([]*models.Country, error) {
	r, err := repositories.DoSimpleQuery(fetchAllStmt)

	if err != nil {
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateCountry)
}
