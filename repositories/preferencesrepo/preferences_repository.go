package preferencesrepo

import (
	"vocaportal/models"
	"vocaportal/repositories"
)

const (
	fetchAllSQL = `
	SELECT
	ID,
	NAME
	FROM preferences
	`

	searchByNameSQL = `
	SELECT
	ID,
	NAME
	FROM preferences
	WHERE NAME LIKE $1
	`
)

// Stmt creator
var stmtCreator = repositories.NewStmtCreator("preferences")

// Stmts
var (
	fecthAllStmt     = stmtCreator.NewStmt(fetchAllSQL)
	searchByNameStmt = stmtCreator.NewStmt(searchByNameSQL)
)

func FetchAll() ([]*models.Preference, error) {
	r, err := repositories.DoSimpleQuery(fecthAllStmt)

	if err != nil {
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreatePreference)
}

func SearchByName(name string) ([]*models.Preference, error) {
	r, err := repositories.DoSimpleQuery(searchByNameStmt, name)

	if err != nil {
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreatePreference)
}
