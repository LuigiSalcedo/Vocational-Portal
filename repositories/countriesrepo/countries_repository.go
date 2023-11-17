package countriesrepo

import (
	"database/sql"
	"vocaportal/core"
	"vocaportal/database"
	"vocaportal/models"
	"vocaportal/repositories"
)

const (
	searchByName = `
	SELECT
	ID, 
	NAME
	FROM countries
	WHERE NAME LIKE $1
	`

	fetchCountry = `
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
	searchByNameStmt = createStmt(searchByName)
	fetchCountryStmt = createStmt(fetchCountry)
)

func SearchByName(name string) ([]models.Country, error) {
	r, err := repositories.DoSimpleQuery(searchByNameStmt, name)

	if err != nil {
		core.LogErr(err)
		return nil, err
	}

	return repositories.RowsToSlice(r,
		models.CreateCountry, models.ExtractCountry, models.RecoveryCountry)
}

func FetchCountry(id int64) (*models.Country, error) {
	r, err := repositories.DoSimpleQuery(fetchCountryStmt, id)

	if err != nil {
		core.LogErr(err)
		return nil, err
	}

	return repositories.Data(r,
		models.CreateCountry,
		models.ExtractCountry,
		models.RecoveryCountry)
}
