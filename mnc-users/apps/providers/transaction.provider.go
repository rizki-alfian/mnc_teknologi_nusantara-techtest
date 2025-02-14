package providers

import (
	"gorm.io/gorm"
	"mnc-users/apps/databases/repositories"
	"mnc-users/apps/transactions"
)

func NewTransactionService(db *gorm.DB, transactionRepo *repositories.TransactionRepository) *transactions.TransactionService {
	return transactions.NewTransactionService(db, transactionRepo)
}