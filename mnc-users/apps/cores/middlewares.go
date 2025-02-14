package cores

import (
	"mnc-users/apps/cores/middlewares"
	"github.com/labstack/echo/v4"
)

func RegisterMiddlewares(e *echo.Echo) {
	middlewares.InitLogger()
	e.Use(middlewares.DefaultLogger())
}