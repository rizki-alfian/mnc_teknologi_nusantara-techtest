package transactions

import (
	"net/http"
	"mnc-users/apps/transactions/dto"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	Service *TransactionService
}

func NewTransactionHandler(service *TransactionService) *TransactionHandler {
	return &TransactionHandler{Service: service}
}

func (h *TransactionHandler) TopUpTransaction(c echo.Context) error {
	userID := c.Get("user_id").(string)
	var req dto.TopUpRequestDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	topUpReq := dto.TransactionRequestDTO{
		Amount:  req.Amount,
	}

	transactionType := "CREDIT"
	response, err := h.Service.CreateTransaction(userID, topUpReq, transactionType)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "SUCCESS",
		"result": response,
	})
}

func (h *TransactionHandler) PaymentTransaction(c echo.Context) error {
	userID := c.Get("user_id").(string)
	var req dto.PaymentRequestDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	paymentReq := dto.TransactionRequestDTO{
		Amount:  req.Amount,
		Remarks: req.Remarks,
	}

	transactionType := "DEBIT"
	response, err := h.Service.CreateTransaction(userID, paymentReq, transactionType)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "SUCCESS",
		"result": response,
	})
}

func (h *TransactionHandler) TransferTransaction(c echo.Context) error {
	userID := c.Get("user_id").(string)
	var req dto.TransactionRequestDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	transactionType := "DEBIT"
	response, err := h.Service.CreateTransaction(userID, req, transactionType)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "SUCCESS",
		"result": response,
	})
}

func (h *TransactionHandler) GetTransactionsReport(c echo.Context) error {
	userID := c.Get("user_id").(string)

	transactions, err := h.Service.GetTransactionsReport(userID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthenticated"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "SUCCESS",
		"result": transactions,
	})
}