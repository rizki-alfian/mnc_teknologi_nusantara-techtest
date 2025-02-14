package routes

import (
	"github.com/labstack/echo/v4"
	"mnc-users/apps/users"
	"mnc-users/apps/cores/middlewares"
)

func UserRoutes(g *echo.Group, h *users.UserHandler) {
	user := g.Group("/users")
	user.POST("/register", h.Register)
	user.POST("/login", h.Login)

	protected := user.Group("")
	protected.Use(middlewares.JWTMiddleware)
	protected.PUT("/profile", h.UpdateProfile)
}