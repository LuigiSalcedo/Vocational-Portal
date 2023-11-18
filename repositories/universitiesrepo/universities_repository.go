package universitiesrepo

import (
	"vocaportal/models"
	"vocaportal/repositories"
)

// SQL Queries
const (
	searchByNameSQL = `
	SELECT
	U.ID,
	U.NAME,
	City.ID,
	City.NAME,
	Country.ID,
	Country.NAME
	FROM
	universities as U 
	JOIN cities as City ON U.CITY_ID = City.ID
	JOIN countries as Country ON City.COUNTRY_ID = Country.ID
	WHERE U.NAME LIKE $1
	`

	fetchUniversitySQL = `
	SELECT
	U.ID,
	U.NAME,
	City.ID,
	City.NAME,
	Country.ID,
	Country.NAME
	FROM
	universities as U 
	JOIN cities as City ON U.CITY_ID = City.ID
	JOIN countries as Country ON City.COUNTRY_ID = Country.ID
	WHERE U.ID = $1 LIMIT 1
	`

	fetchAllSQL = `
	SELECT
	U.ID,
	U.NAME,
	City.ID,
	City.NAME,
	Country.ID,
	Country.NAME
	FROM
	universities as U 
	JOIN cities as City ON U.CITY_ID = City.ID
	JOIN countries as Country ON City.COUNTRY_ID = Country.ID
	`
)

// Stmt creator
var stmtCreator = repositories.NewStmtCreator("universities")

// Stmts
var (
	searchByNameStmt    = stmtCreator.NewStmt(searchByNameSQL)
	fetchUniversityStmt = stmtCreator.NewStmt(fetchUniversitySQL)
	fetchAllStmt        = stmtCreator.NewStmt(fetchAllSQL)
)

// Return list of universities searched by name
func SearchByName(name string) ([]*models.University, error) {
	r, err := repositories.DoSimpleQuery(searchByNameStmt, name)

	if err != nil {
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateUniversity)
}

// Fetch an university from the database using the ID
func FetchUniversity(id int64) (*models.University, error) {
	r, err := repositories.DoSimpleQuery(fetchUniversityStmt, id)

	if err != nil {
		return nil, err
	}

	result, err := repositories.Data(r, models.CreateUniversity)

	if result == nil {
		return nil, err
	}

	return *result, err
}

// Fetch all universities register
func FetchAll() ([]*models.University, error) {
	r, err := repositories.DoSimpleQuery(fetchAllStmt)

	if err != nil {
		return nil, err
	}

	return repositories.RowsToSlice(r, models.CreateUniversity)
}
