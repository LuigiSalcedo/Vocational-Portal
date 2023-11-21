package areasrepo

import (
	"vocaportal/models"
	"vocaportal/repositories"
)

// SQL Queris
const (
	fetchAllSQL = `
	SELECT
	ID,
	NAME
	FROM study_areas
	`

	searchByNameSQL = `
	SELECT
	ID,
	NAME
	FROM study_areas
	WHERE NAME LIKE $1
	`
)

// Stmt creator
var stmtCreator = repositories.NewStmtCreator("study areas")

// Stmts
var (
	fetchAllStmt     = stmtCreator.NewStmt(fetchAllSQL)
	searchByNameStmt = stmtCreator.NewStmt(searchByNameSQL)
)

func FetchAll() ([]*models.Area, error) {
	r, err := repositories.DoSimpleQuery(fetchAllStmt)

	if err != nil {
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateArea)
}

func SearchByName(name string) ([]*models.Area, error) {
	r, err := repositories.DoSimpleQuery(searchByNameStmt, name)

	if err != nil {
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateArea)
}
