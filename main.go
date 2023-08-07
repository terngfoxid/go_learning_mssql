package main

import (
	"go-mssql/Config"
	"go-mssql/Routes"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	//Start DB Connection
	Config.StartConnection()
	defer Config.DB.Close()

	r := Routes.SetupRouter()
	//running
	r.Run(":8080")
}
