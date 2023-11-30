package adminsrepo

import (
	"vocaportal/core"
	"vocaportal/repositories"
)

const (
	insertSQL = "INSERT INTO ADMINS VALUES($1, $2)"
	fetchSQL  = "SELECT PASSWORD FROM ADMINS WHERE name = $1"
)

// Stmt creator
var stmtCreator = repositories.NewStmtCreator("admins")

// Stmt
var (
	insertStmt = stmtCreator.NewStmt(insertSQL)
	fetchStmt  = stmtCreator.NewStmt(fetchSQL)
)

func Save(name, password string) error {
	_, err := insertStmt.Exec(name, password)

	if err != nil {
		core.LogErr(err)
		return err
	}

	return nil
}

func Fetch(name string) (string, error) {
	r, err := repositories.DoSimpleQuery(fetchStmt)

	if err != nil {
		return "", err
	}

	if !r.Next() {
		return "", nil
	}

	str := ""

	err = r.Scan(&str)

	return str, err
}
