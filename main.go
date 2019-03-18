package main

import (
	"./Database"
	"./Route"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	Database.ConfigDatabase()
	Route.ConfigRoute()
}
