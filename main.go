package main

import (
	//"fortune-cookies/config"
	"fortune-cookies/routes"
	"os"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func main() {
	ech := echo.New()

	//Use route
	routes.InitRoute(ech)

	//Config
	//cfg, _ := config.NewConfig(".env")

	//use CORS
	ech.Use(middleware.CORS())
	
	ech.Logger.Fatal(ech.Start(":" + os.Getenv("PORT")))
}