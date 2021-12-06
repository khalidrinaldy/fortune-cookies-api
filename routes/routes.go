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
	e.GET("/user/count", repository.CountUsers(database))

	//Cart API
	e.GET("/cart", repository.GetCartList(database), middlewares.IsLoggedIn())
	e.POST("/cart", repository.AddToCart(database), middlewares.IsLoggedIn())
	e.PUT("/cart", repository.UpdateItemCart(database), middlewares.IsLoggedIn())
	e.DELETE("/cart/:product_id", repository.DeleteItemCart(database), middlewares.IsLoggedIn())

	//Product API
	e.GET("/product", repository.GetAllProducts(database))
	e.GET("/product/:category", repository.GetProductsByCategory(database))
	e.GET("/product/count", repository.CountProducts(database))
	e.POST("/product", repository.AddProduct(database),middlewares.IsLoggedIn())
	e.PUT("/product/edit/:id", repository.UpdateProduct(database),middlewares.IsLoggedIn())
	e.DELETE("/product/delete/:id", repository.DeleteProduct(database), middlewares.IsLoggedIn())

	//History API
	e.GET("/history", repository.GetAllHistory(database), middlewares.IsLoggedIn())
	e.GET("/history/detail/:history_id", repository.GetDetailHistoryProducts(database), middlewares.IsLoggedIn())
	e.GET("/history/count", repository.CountHistory(database))
	e.POST("/purchase", repository.Purchase(database), middlewares.IsLoggedIn())

	//Admin API
	e.GET("/adminbytoken", repository.GetAdminByToken(database), middlewares.IsLoggedIn())
	e.POST("/admin/register", repository.RegisterAdmin(database))
	e.POST("/admin/login", repository.LoginAdmin(database))
}
