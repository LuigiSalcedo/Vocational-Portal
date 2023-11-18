package programmsrepo

import (
	"vocaportal/models"
	"vocaportal/repositories"
)

// SQL Queries
const (
	searchByNameSQL = `
	SELECT
	ID,
	NAME
	FROM academic_programmes
	WHERE NAME LIKE $1
	`

	fetchProgrammSQL = `
	SELECT
	ID,
	NAME
	FROM academic_programmes
	WHERE ID = $1
	`
)

// Stmt creator
var stmtCreator = repositories.NewStmtCreator("academic programmes")

// Stmts
var (
	searchByNameStmt  = stmtCreator.NewStmt(searchByNameSQL)
	fetchProgrammStmt = stmtCreator.NewStmt(fetchProgrammSQL)
)

func SearchByName(name string) ([]*models.Programm, error) {
	r, err := repositories.DoSimpleQuery(searchByNameStmt, name)

	if err != nil {
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateProgramm)
}

func FetchProgramm(id int64) (*models.Programm, error) {
	r, err := repositories.DoSimpleQuery(fetchProgrammStmt, id)

	if err != nil {
		return nil, err
	}

	p, err := repositories.Data(r, models.CreateProgramm)

	if err != nil {
		return nil, err
	}

	return *p, nil
}
