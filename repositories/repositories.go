package repositories

import (
	"database/sql"
	"vocaportal/core"
)

// Execute a simple query on the database using a stmt
func DoSimpleQuery(stmt *sql.Stmt, args ...any) (*sql.Rows, error) {
	return stmt.Query(args...)
}

// Convert *sql.Rows to a slice
// cf = Creation Func, this func should create register the type T
// ef = Extraction Func, this func should return a slice of pointer to values to get
// rf = Recovery Func, this func should put the readed values to the register
// from de row
func RowsToSlice[T any](r *sql.Rows, cf func() T, ef func(value T) []any, rf func(v *T, args ...any)) ([]T, error) {
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
