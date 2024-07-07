package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTWithConfig(config middleware.JWTConfig) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(config)
}
