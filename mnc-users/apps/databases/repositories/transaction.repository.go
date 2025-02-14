package repositories

import (
	"mnc-users/apps/databases/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(tx *gorm.DB, transaction *models.Transaction) error {
	return tx.Create(transaction).Error
}

func (r *TransactionRepository) GetUserBalance(userID string) (int64, error) {
	var balance int64
	err := r.db.Raw(`
		SELECT COALESCE(SUM(
			CASE 
				WHEN transaction_type = 'CREDIT' THEN amount 
				WHEN transaction_type = 'DEBIT' THEN -amount 
				ELSE 0 
			END
		), 0) FROM transactions WHERE user_id = ?`, userID).Scan(&balance).Error
	return balance, err
}

func (r *TransactionRepository) GetUserTransactions(userID string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("user_id = ?", userID).Order("created_date DESC").Find(&transactions).Error
	return transactions, err
}
