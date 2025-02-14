package container

import (
	"mnc-users/config"
	"mnc-users/apps/providers"
	"mnc-users/apps/databases/repositories"
	"mnc-users/apps/users"
	"mnc-users/apps/test1"
	"mnc-users/apps/transactions"
)

type Container struct {
	UserHandler *users.UserHandler
	Test1Handler *test1.Test1Handler
	TransactionHandler *transactions.TransactionHandler
}

func NewContainer() *Container {
	db := config.InitDB()

	userRepo := repositories.NewUserRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	userService := providers.NewUserService(userRepo)
	test1Service := providers.NewTest1Service()
	transactionService := providers.NewTransactionService(db, transactionRepo)

	userHandler := users.NewUserHandler(userService)
	test1Handler := test1.NewTest1Handler(test1Service)
	transactionHandler := transactions.NewTransactionHandler(transactionService)

	return &Container{
		UserHandler: userHandler,
		Test1Handler: test1Handler,
		TransactionHandler: transactionHandler,
	}
}