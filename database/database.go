package database

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	dbData := os.Getenv("vocaportal_db_data")

	if len(dbData) == 0 {
		log.Fatal("database data not found")
	}

	database, err := sql.Open("postgres", dbData)

	if err != nil {
		log.Fatal(err)
	}

	db = database
}

func DB() *sql.DB {
	return db
}

func Begin() (*sql.Tx, error) {
	return db.Begin()
}

func BeginTx(ctx context.Context) (*sql.Tx, error) {
	return db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
}

func Conn(ctx context.Context) (*sql.Conn, error) {
	return db.Conn(ctx)
}

func InitStmt(sql, repository string) *sql.Stmt {
	stmt, err := db.Prepare(sql)

	if err != nil {
		log.Fatalf("Error in '%s' repository: %v\n", repository, err)
	}

	return stmt
}
