package repositories

import (
	"database/sql"
	"vocaportal/core"
)

// A creation function is a func that should create a struct register
type CreationFunction[T any] func() T

// Extractor function is a func that should return the atributtes of a register
type ExtractorFunc[T any] func(strct T) []any

// Recovery function is a func that should put variable on a struct register
type RecoveryFunc[T any] func(strct *T, data ...any)

// Execute a simple query on the database using a stmt
func DoSimpleQuery(stmt *sql.Stmt, args ...any) (*sql.Rows, error) {
	return stmt.Query(args...)
}

// Convert *sql.Rows to a slice
// cf = Creation Func, this func should create register the type T
// ef = Extraction Func, this func should return a slice of pointer to values to get
// rf = Recovery Func, this func should put the readed values to the register
// from de row
func RowsToSlice[T any](r *sql.Rows, cf CreationFunction[T], ef ExtractorFunc[T], rf RecoveryFunc[T]) ([]T, error) {
	result := make([]T, 0, 20)

	for r.Next() {
		v := cf()
		data := ef(v)
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

		rf(&v, data...)

		result = append(result, v)
	}

	return result, nil
}

// Get the first register by a sql query. This function only should be used when
// the query going to return 1 or 0 register. If this function return nil, nil so
// the register was NotFound.
func Data[T any](r *sql.Rows, cf CreationFunction[T], ef ExtractorFunc[T], rf RecoveryFunc[T]) (*T, error) {
	slc, err := RowsToSlice(r, cf, ef, rf)

	if err != nil {
		return nil, err
	}

	if len(slc) == 0 {
		return nil, nil
	}

	return &slc[0], nil
}
