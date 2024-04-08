package controllers

import (
	"context"
	"time"

	"github.com/asynched/idempotent-transaction-api/domain/entities"
	"github.com/asynched/idempotent-transaction-api/domain/repositories"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type TransactionController struct {
	accountRepository     *repositories.AccountRepository
	transactionRepository *repositories.TransactionRepository
	redis                 *redis.Client
}

func NewTransactionController(
	accountRepository *repositories.AccountRepository,
	transactionRepository *repositories.TransactionRepository,
	redis *redis.Client,
) *TransactionController {
	return &TransactionController{
		accountRepository,
		transactionRepository,
		redis,
	}
}

type CreateTransactionDto struct {
	Amount float64 `json:"amount"`
	Payer  string  `json:"payer"`
	Payee  string  `json:"payee"`
}

func (ctrl *TransactionController) ListAll(c echo.Context) error {
	transactions, err := ctrl.transactionRepository.FindAllById(
		c.Param("id"),
	)

	if err != nil {
		return c.JSON(500, echo.Map{
			"message": "Failed to list transactions",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, transactions)
}

func (ctrl *TransactionController) Create(c echo.Context) error {
	idempotencyKey := c.Request().Header.Get("X-Idempotency-Key")

	if idempotencyKey == "" {
		return c.JSON(400, echo.Map{
			"message": "Idempotency key is required",
			"error":   "X-Idempotency-Key header is missing",
		})
	}

	if ctrl.redis.Exists(context.Background(), idempotencyKey).Val() == 1 {
		return c.JSON(200, echo.Map{
			"message": "Transaction already processed",
		})
	}

	data := CreateTransactionDto{}

	if err := c.Bind(&data); err != nil {
		return c.JSON(400, echo.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	payer, err := ctrl.accountRepository.FindById(data.Payer)

	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "Payer not found",
			"error":   err.Error(),
		})
	}

	_, err = ctrl.accountRepository.FindById(data.Payee)

	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "Payee not found",
			"error":   err.Error(),
		})
	}

	if payer.Balance < data.Amount {
		return c.JSON(400, echo.Map{
			"message": "Insufficient balance",
		})
	}

	transaction, err := ctrl.transactionRepository.Create(entities.Transaction{
		Amount: data.Amount,
		Payer:  data.Payer,
		Payee:  data.Payee,
	})

	if err != nil {
		return c.JSON(500, echo.Map{
			"message": "Failed to create transaction",
			"error":   err.Error(),
		})
	}

	ctrl.redis.Set(context.Background(), idempotencyKey, "1", time.Minute*10)

	return c.JSON(201, transaction)
}
