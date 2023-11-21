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
)

// Stmt creator
var stmtCreator = repositories.NewStmtCreator("study areas")

// Stmts
var (
	fetchAllStmt = stmtCreator.NewStmt(fetchAllSQL)
)

func FetchAll() ([]*models.Area, error) {
	r, err := repositories.DoSimpleQuery(fetchAllStmt)

	if err != nil {
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateArea)
}
