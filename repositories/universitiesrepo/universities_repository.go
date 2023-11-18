package universitiesrepo

import (
	"vocaportal/core"
	"vocaportal/models"
	"vocaportal/repositories"
)

// SQL Queries
const (
	searchByNameSQL = `
	SELECT
	ID,
	NAME
	FROM UNIVERSITIES
	WHERE NAME LIKE $1
	`

	fetchUniversitySQL = `
	SELECT
	ID,
	NAME
	FROM universities
	WHERE ID = $1 LIMIT 1
	`
)

// Stmt creator
var stmtCreator = repositories.NewStmtCreator("universities")

// Stmts
var (
	searchByNameStmt    = stmtCreator.NewStmt(searchByNameSQL)
	fetchUniversityStmt = stmtCreator.NewStmt(fetchUniversitySQL)
)

// Return list of universities searched by name
func SearchByName(name string) ([]*models.University, error) {
	r, err := repositories.DoSimpleQuery(searchByNameStmt, name)

	if err != nil {
		core.LogErr(err)
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateUniversity)
}

// Fetch an university from the database using the ID
func FetchUniversity(id int64) (*models.University, error) {
	r, err := repositories.DoSimpleQuery(fetchUniversityStmt, id)

	if err != nil {
		core.LogErr(err)
		return nil, err
	}

	result, err := repositories.Data(r, models.CreateUniversity)

	if result == nil {
		return nil, nil
	}

	return *result, err
}
