package repositories

import (
	"database/sql"
	"vocaportal/core"
	"vocaportal/database"
	"vocaportal/models"
)

// A StmtCreator is a type that create a Stmt for a specific repository
type stmtCreator struct {
	repositoryName string
}

// A creation function is a func that should create a struct register
type CreationFunction[T models.DatabaseEntity] func() T

// Create a new StmtCreator using the name
func NewStmtCreator(reponame string) stmtCreator {
	return stmtCreator{
		repositoryName: reponame,
	}
}

// Create a new stmt
func (sc stmtCreator) NewStmt(sql string) *sql.Stmt {
	return database.InitStmt(sql, sc.repositoryName)
}

// Execute a simple query on the database using a stmt
func DoSimpleQuery(stmt *sql.Stmt, args ...any) (*sql.Rows, error) {
	return stmt.Query(args...)
}

// Convert *sql.Rows to a slice
// cf = Creation Func, this func should create register the type T
// ef = Extraction Func, this func should return a slice of pointer to values to get
// rf = Recovery Func, this func should put the readed values to the register
// from de row
func RowsToSlice[T models.DatabaseEntity](r *sql.Rows, cf CreationFunction[T]) ([]T, error) {
	result := make([]T, 0, 20)

	for r.Next() {
		v := cf()
		data := v.Extract()
		scdata := make([]any, len(data))

		for i := range data {
			scdata[i] = &data[i]
		}

		err := r.Scan(scdata...)

		if err != nil {
			core.LogErr(err)
			return nil, err
		}

		for i := range data {
			data[i] = *scdata[i].(*any)
		}

		v.Recovery(data...)

		result = append(result, v)
	}

	return result, nil
}

// Get the first register by a sql query. This function only should be used when
// the query going to return 1 or 0 register. If this function return nil, nil so
// the register was NotFound.
func Data[T models.DatabaseEntity](r *sql.Rows, cf CreationFunction[T]) (*T, error) {
	slc, err := RowsToSlice(r, cf)

	if err != nil {
		core.LogErr(err)
		return nil, err
	}

	if len(slc) == 0 {
		return nil, nil
	}

	return &slc[0], nil
}
