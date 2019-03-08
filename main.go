package main

import (
	"github.com/tkitsunai/minimum-clean-architecture/config"
	"github.com/tkitsunai/minimum-clean-architecture/db"
	"github.com/tkitsunai/minimum-clean-architecture/rest"
)

func main() {
	con := db.New(db.DialectMysql, config.DBConnection().Source)
	defer con.CloseDB()
	rest.NewApplication(
		con.GetDB(),
	).Run()
}
