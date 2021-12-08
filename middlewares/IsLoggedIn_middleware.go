package middlewares

import (
	"os"

	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

func IsLoggedIn() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	})
}