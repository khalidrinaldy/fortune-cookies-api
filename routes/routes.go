package routes

import (
	"fortune-cookies/db"
	"fortune-cookies/middlewares"
	"fortune-cookies/repository"

	"github.com/labstack/echo"
)

func InitRoute(e *echo.Echo) {
	//Init db
	database := db.OpenDatabase()

	//Call Migrate Func
	db.Migrate(database)

	//Authentication API
	e.POST("/register", repository.Registration(database))
	e.POST("/login", repository.Login(database))
	e.GET("/user/:id", repository.GetUser(database))
	e.GET("/userbytoken", repository.GetUserByToken(database), middlewares.IsLoggedIn())

	//Cart API
	e.GET("/cart/:id", repository.GetCartList(database), middlewares.IsLoggedIn())
	e.POST("/cart", repository.AddToCart(database), middlewares.IsLoggedIn())
	e.PUT("/cart/:id", repository.UpdateItemCart(database), middlewares.IsLoggedIn())
	e.DELETE("/cart/:id/:product_id", repository.DeleteItemCart(database), middlewares.IsLoggedIn())

	//Product API
	e.GET("/product", repository.GetAllProducts(database))
	e.GET("/product/:category", repository.GetProductsByCategory(database))
	e.POST("/product", repository.AddProduct(database))
	e.PUT("/product/:id", repository.UpdateProduct(database))
	e.DELETE("/product/:id", repository.DeleteProduct(database))
}
