package routes

import (
	"github.com/labstack/echo/v4"
	"mnc-users/apps/transactions"
	"mnc-users/apps/cores/middlewares"
)

func TransactionRoutes(g *echo.Group, h *transactions.TransactionHandler) {
	transaction := g.Group("/transactions")
	protected := transaction.Group("")
	protected.Use(middlewares.JWTMiddleware)
	protected.POST("/topup", h.TopUpTransaction)
	protected.POST("/pay", h.PaymentTransaction)
	protected.POST("/transfer", h.TransferTransaction)
	protected.GET("/all", h.GetTransactionsReport)
}