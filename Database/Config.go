package Database

import (
	"berita/Helper"
	"database/sql"
)

func ConfigDatabase() *sql.DB {
	var (
		err      error
		database *sql.DB
	)
	database, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/berita")
	Helper.LogError(err)
	return database
}
