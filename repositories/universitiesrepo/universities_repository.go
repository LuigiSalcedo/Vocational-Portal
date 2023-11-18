package universitiesrepo

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

// Create a stmt in universities repository
func createStmt(sql string) *sql.Stmt {
	return database.InitStmt(sql, "universities")
}

// Prepared Statements
var (
	searchByNameStmt    = createStmt(searchByNameSQL)
	fetchUniversityStmt = createStmt(fetchUniversitySQL)
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

	return *result, err
}
