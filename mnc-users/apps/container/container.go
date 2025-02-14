package container

import (
	"mnc-users/config"
	"mnc-users/apps/providers"
	"mnc-users/apps/databases/repositories"
	"mnc-users/apps/users"
	"mnc-users/apps/test1"
)

type Container struct {
	UserHandler *users.UserHandler
	Test1Handler *test1.Test1Handler
}

func NewContainer() *Container {
	db := config.InitDB()

	userRepo := repositories.NewUserRepository(db)

	userService := providers.NewUserService(userRepo)
	test1Service := providers.NewTest1Service()

	userHandler := users.NewUserHandler(userService)
	test1Handler := test1.NewTest1Handler(test1Service)

	return &Container{
		UserHandler: userHandler,
		Test1Handler: test1Handler,
	}
}