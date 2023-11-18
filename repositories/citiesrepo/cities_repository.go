package citiesrepo

import (
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories"
)

// SQL Queries
const (
	searchByNameSQL = `
	SELECT
	CITIES.ID,
	CITIES.NAME,
	COUNTRIES.ID,
	COUNTRIES.NAME
	FROM cities JOIN countries ON cities.COUNTRY_ID = countries.ID
	WHERE cities.NAME LIKE $1
	`
)

// Stmt creator
var stmtCreator = repositories.NewStmtCreator("cities")

// Statements
var (
	searchByNameStmt = stmtCreator.NewStmt(searchByNameSQL)
)

// Get a list of cities using the name
func SearchByName(name string) ([]*models.City, error) {
	r, err := repositories.DoSimpleQuery(searchByNameStmt, name)

	if err != nil {
		core.LogErr(err)
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateCity)
}
