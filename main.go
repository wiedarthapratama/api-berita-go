package main

import (
	"berita/Database"
	"berita/Route"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	Database.ConfigDatabase()
	Route.ConfigRoute()
}
