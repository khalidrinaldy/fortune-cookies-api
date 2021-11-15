package middlewares

import (
	"fortune-cookies/config"

	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

func IsLoggedIn() echo.MiddlewareFunc {
	cfg, _ := config.NewConfig(".env")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(cfg.JWTConfig.SecretKey),
	})
}