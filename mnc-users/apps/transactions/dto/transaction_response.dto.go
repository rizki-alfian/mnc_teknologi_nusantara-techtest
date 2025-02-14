package dto

type TopUpResponseDTO struct {
    TopUpID      string `json:"top_up_id"`
    AmountTopUp   string `json:"amount_top_up"`
    BalanceBefore    string `json:"balance_before"`
    BalanceAfter string `json:"balance_after"`
    CreatedDate string `json:"created_date"`
}