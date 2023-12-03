package programmsrepo

import (
	"fmt"
	"strings"
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
	NAME,
	DESCRIPTION
	FROM academic_programmes
	WHERE NAME LIKE $1
	`

	fetchProgrammSQL = `
	SELECT
	ID,
	NAME,
	DESCRIPTION
	FROM academic_programmes
	WHERE ID = $1
	`

	fetchAllSQL = `
	SELECT
	ID,
	NAME,
	DESCRIPTION
	FROM academic_programmes
	`

	searchByAreaRelation = `
	SELECT
	P.ID,
	P.NAME,
	P.DESCRIPTION,
	SCORE
	FROM programmes_areas as A JOIN academic_programmes as P ON P.ID = A.programm_id
	WHERE 1 = 1
	`
)

// Stmt creator
var stmtCreator = repositories.NewStmtCreator("academic programmes")

// Stmts
var (
	searchByNameStmt  = stmtCreator.NewStmt(searchByNameSQL)
	fetchProgrammStmt = stmtCreator.NewStmt(fetchProgrammSQL)
	fetchAllStmt      = stmtCreator.NewStmt(fetchAllSQL)
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

func FetchAll() ([]*models.Programm, error) {
	r, err := repositories.DoSimpleQuery(fetchAllStmt)

	if err != nil {
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateProgramm)
}

func SearchByAreaRelation(areasIds []int64) ([]*models.PWAR, error) {
	subQuery := strings.Builder{}

	initialQuery := searchByAreaRelation

	subQuery.WriteString(initialQuery)

	if len(areasIds) >= 1 {
		subQuery.WriteString("\nAND (")
	}

	for i := range areasIds {
		subQuery.WriteString(fmt.Sprintf("\nA.area_id = %d", areasIds[i]))
		if i < len(areasIds)-1 {
			subQuery.WriteString("\nOR")
		} else {
			subQuery.WriteString(")")
		}
	}

	finalQuery := `
	SELECT
	ID,
	NAME,
	DESCRIPTION,
	AVG(SCORE) as N
	FROM (
	` + subQuery.String() + ") as sq GROUP BY(ID, NAME, DESCRIPTION) ORDER BY(N) DESC"

	r, err := database.DB().Query(finalQuery)

	if err != nil {
		core.LogErr(err)
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreatePWAR)
}
