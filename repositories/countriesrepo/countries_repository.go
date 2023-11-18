package countriesrepo

import (
	"database/sql"
	"vocaportal/core"
	"vocaportal/database"
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
)

func createStmt(sql string) *sql.Stmt {
	return database.InitStmt(sql, "countries")
}

var (
	searchByNameStmt = createStmt(searchByNameSQL)
	fetchCountryStmt = createStmt(fetchCountrySQL)
)

func SearchByName(name string) ([]*models.Country, error) {
	r, err := repositories.DoSimpleQuery(searchByNameStmt, name)

	if err != nil {
		core.LogErr(err)
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateCountry)
}

func FetchCountry(id int64) (*models.Country, error) {
	r, err := repositories.DoSimpleQuery(fetchCountryStmt, id)

	if err != nil {
		core.LogErr(err)
		return nil, err
	}

	result, err := repositories.Data(r, models.CreateCountry)

	return *result, err
}
