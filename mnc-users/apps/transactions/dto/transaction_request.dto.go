package dto

type TopUpRequestDTO struct {
	Amount   int64 `json:"amount,omitempty"`
}

type PaymentRequestDTO struct {
	Amount   int64 `json:"amount"`
	Remarks   string `json:"remarks"`
}

type TransactionRequestDTO struct {
	TargetUser   *string `json:"target_user,omitempty"`
	Amount   int64 `json:"amount"`
	Remarks   string `json:"remarks,omitempty"`
}