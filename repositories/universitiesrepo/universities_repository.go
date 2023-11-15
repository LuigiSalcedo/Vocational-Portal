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
)

// Create a stmt in universities repository
func createStmt(sql string) *sql.Stmt {
	return database.InitStmt(sql, "universities")
}

// Prepared Statements
var (
	searchByNameStmt = createStmt(searchByNameSQL)
)

func SearchByName(name string) ([]models.University, error) {
	r, err := repositories.DoSimpleQuery(searchByNameStmt, name)

	if err != nil {
		core.LogErr(err)
		return nil, err
	}

	return repositories.RowsToSlice(r,
		models.CreateUniversity,
		models.ExtractUniversity,
		models.RecoveryUniversity)
}
