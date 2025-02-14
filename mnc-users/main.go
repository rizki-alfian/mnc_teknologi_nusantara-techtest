package main

import (
	"github.com/joho/godotenv"
	"os"
	"log"

	"mnc-users/apps/cores"
	"mnc-users/apps/container"
	"mnc-users/apps/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
        log.Fatal("Warning: .env file not found, using system environment variables")
    }

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	e := echo.New()

	di := container.NewContainer()

	cores.RegisterMiddlewares(e)

	routes.SetupRoutes(e, di)

	APP_PORT := os.Getenv("APP_PORT")
	if APP_PORT == "" {
        log.Fatal("APP_PORT not set in environment variables")
    }
	e.Start(":" + APP_PORT)
}