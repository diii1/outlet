package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type LoggerConfig struct {
	Skipper          interface{}
	Format           string
	CustomTimeFormat string
}

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper:          nil,
		Format:           "time=${time_rfc3339},  method=${method}, uri=${host}${path}, status=${status}, latency_human=${latency_human} \n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           nil,
	}))
}
