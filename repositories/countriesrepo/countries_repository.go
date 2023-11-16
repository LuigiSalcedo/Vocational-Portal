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
)

func createStmt(sql string) *sql.Stmt {
	return database.InitStmt(sql, "countries")
}

var (
	searchByNameStmt = createStmt(searchByName)
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
