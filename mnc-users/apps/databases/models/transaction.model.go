package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID   uuid.UUID  `json:"transaction_id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID          uuid.UUID  `json:"user_id"`
	TargetUser      *uuid.UUID `json:"target_user,omitempty"`
	TransactionType string     `json:"transaction_type"`
	Amount          int64      `json:"amount"`
	Remarks         *string    `json:"remarks,omitempty"`
	BalanceBefore   int64      `json:"balance_before"`
	BalanceAfter    int64      `json:"balance_after"`
	CreatedDate     time.Time  `json:"created_date" gorm:"default:now()"`
}
