package main

import (
	"prakerja11/configs"
	"prakerja11/routes"

	"github.com/labstack/echo/v4"
)


func main(){
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoutes(e)
	e.Start(":8000")
}


