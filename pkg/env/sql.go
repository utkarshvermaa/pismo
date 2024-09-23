package env

import (
	"database/sql"
	"os"
)

func NewDB() *sql.DB {
	sqlDriver := os.Getenv("SQL_DRIVER")
	sqlDatasource := os.Getenv("SQL_DATA_SRC")
	_db, err := sql.Open(sqlDriver, sqlDatasource)
	if err != nil {
		panic("error while opening sql connection: " + err.Error())
	}

	if err := _db.Ping(); err != nil {
		panic("sql driver " + sqlDriver + " sqlDatasource " + sqlDatasource + " error while pinging sql connection: " + err.Error())
	}
	return _db
}
