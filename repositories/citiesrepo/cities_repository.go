package citiesrepo

import (
	"database/sql"
	"vocaportal/core"
	"vocaportal/database"
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
func createStmt(sql string) *sql.Stmt {
	return database.InitStmt(sql, "cities")
}

// Statements
var (
	searchByNameStmt = createStmt(searchByNameSQL)
)

// Get a list of cities using the name
func SearchByName(name string) ([]models.City, error) {
	r, err := repositories.DoSimpleQuery(searchByNameStmt, name)

	if err != nil {
		core.LogErr(err)
		return nil, err
	}

	return repositories.RowsToSlice(r,
		models.CreateCity,
		models.ExtractCity,
		models.RecoveryCity)
}
