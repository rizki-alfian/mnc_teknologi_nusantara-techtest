package transactions

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"mnc-users/apps/databases/models"
	"mnc-users/apps/databases/repositories"
	"mnc-users/apps/transactions/dto"
	"fmt"
)

type TransactionService struct {
	db   *gorm.DB
	repo *repositories.TransactionRepository
}

func NewTransactionService(db *gorm.DB, repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{
		db:   db,
		repo: repo,
	}
}

type TransactionResponse struct {
	TransactionID   uuid.UUID `json:"transaction_id"`
	Amount          int64     `json:"amount"`
	Remarks         string    `json:"remarks"`
	BalanceBefore   int64     `json:"balance_before"`
	BalanceAfter    int64     `json:"balance_after"`
	CreatedDate     string    `json:"created_date"`
	TransactionType string    `json:"transaction_type"`
}

func (s *TransactionService) CreateTransaction(userID string, req dto.TransactionRequestDTO, transactionType string) (*TransactionResponse, error) {
	if req.Amount <= 0 {
		return nil, errors.New("Invalid transaction amount")
	}

	var response *TransactionResponse
	err := s.db.Transaction(func(tx *gorm.DB) error {
		balanceBefore, err := s.repo.GetUserBalance(userID)
		if err != nil {
			return err
		}

		if transactionType == "DEBIT" && balanceBefore < req.Amount {
			return errors.New("Balance is not enough")
		}

		balanceAfter := balanceBefore + req.Amount
		if transactionType == "DEBIT" {
			balanceAfter = balanceBefore - req.Amount
		}

		transaction := &models.Transaction{
			TransactionID:   uuid.New(),
			UserID:          uuid.MustParse(userID),
			TransactionType: transactionType,
			Amount:          req.Amount,
			Remarks:         &req.Remarks,
			BalanceBefore:   balanceBefore,
			BalanceAfter:    balanceAfter,
			CreatedDate:     time.Now(),
		}

		if req.TargetUser != nil {
			targetUUID := uuid.MustParse(*req.TargetUser)
			transaction.TargetUser = &targetUUID
		}

		if err := s.repo.CreateTransaction(tx, transaction); err != nil {
			return err
		}

		response = &TransactionResponse{
			TransactionID:   transaction.TransactionID,
			Amount:          transaction.Amount,
			Remarks:         *transaction.Remarks,
			BalanceBefore:   transaction.BalanceBefore,
			BalanceAfter:    transaction.BalanceAfter,
			CreatedDate:     transaction.CreatedDate.Format("2006-01-02 15:04:05"),
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *TransactionService) GetTransactionsReport(userID string) ([]TransactionResponse, error) {
	fmt.Println("userID,", userID)
	transactions, err := s.repo.GetUserTransactions(userID)

	if err != nil {
		return nil, err
	}

	var reports []TransactionResponse
	for _, t := range transactions {
		reports = append(reports, TransactionResponse{
			TransactionID:   t.TransactionID,
			Amount:          t.Amount,
			Remarks:         *t.Remarks,
			BalanceBefore:   t.BalanceBefore,
			BalanceAfter:    t.BalanceAfter,
			CreatedDate:     t.CreatedDate.Format("2006-01-02 15:04:05"),
			TransactionType: t.TransactionType,
		})
	}

	return reports, nil
}