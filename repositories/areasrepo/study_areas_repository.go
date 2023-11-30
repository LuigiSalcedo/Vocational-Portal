package areasrepo

import (
	"fmt"
	"strings"
	"vocaportal/core"
	"vocaportal/database"
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

	searchByPreferencesSQL = `
	SELECT
	DISTINCT
	A.ID
	FROM
	study_areas as A 
	JOIN preferences_areas as PA ON A.ID = PA.Area_ID
	JOIN preferences as P on P.ID = PA.Preference_ID
	WHERE 1 = 1
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

func SearchByPreference(preferencesIds []int64) ([]int64, error) {
	query := strings.Builder{}

	query.WriteString(searchByPreferencesSQL)

	if len(preferencesIds) >= 1 {
		query.WriteString("\nAND (")
	}

	for i := range preferencesIds {
		query.WriteString(fmt.Sprintf("\nPA.Preference_ID = %d", preferencesIds[i]))
		if i < len(preferencesIds)-1 {
			query.WriteString("\nOR")
		} else {
			query.WriteString(")")
		}
	}

	r, err := database.DB().Query(query.String())

	if err != nil {
		core.LogErr(err)
		return nil, err
	}

	var a int64
	result := make([]int64, 0, 10)

	for r.Next() {
		err = r.Scan(&a)

		if err != nil {
			core.LogErr(err)
			return nil, err
		}

		result = append(result, a)
	}

	return result, nil
}
