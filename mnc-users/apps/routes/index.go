package routes

import (
	"github.com/labstack/echo/v4"
	"mnc-users/apps/container"
)

func SetupRoutes(e *echo.Echo, di *container.Container) {
	api := e.Group("/api")

	Test1Routes(api, di.Test1Handler)
	UserRoutes(api, di.UserHandler)
}