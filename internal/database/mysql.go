package database

import (
	"database/sql"

	"github.com/cleitonSilvaViana/social-go/config"

	_ "github.com/go-sql-driver/mysql"
)


func ConnectToMysql() (*sql.DB, error) {

	db, err := sql.Open("mysql", config.STRING_MYSQL_CONNECTION)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}