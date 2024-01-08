package main

import (
	"firly/mytaskapp/apps/config"
	"firly/mytaskapp/apps/database"
	"firly/mytaskapp/apps/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// logging := helpers.NewLogger()
	cfg := config.InitConfig()
	dbMysql := database.InitDBMysql(cfg)

	//call migration
	database.InitialMigration(dbMysql)

	//create a new echo instance
	e := echo.New()
	e.Use(middleware.CORS())
	//remove pre trailingslash
	e.Pre(middleware.RemoveTrailingSlash())

	//e.Use middleware logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.InitRouter(dbMysql, e)

	//start server and port
	e.Logger.Fatal(e.Start(":8083"))
}
