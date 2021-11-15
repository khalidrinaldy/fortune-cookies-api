package main

import (
	"fmt"
	"fortune-cookies/config"
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
	cfg, _ := config.NewConfig(".env")

	//use CORS
	ech.Use(middleware.CORS())
	
	//Set PORT
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("port no ERROR")
		port = cfg.Port
	}

	ech.Logger.Fatal(ech.Start(":" + port))
}