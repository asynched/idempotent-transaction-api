package repositories

import (
	"database/sql"

	"github.com/asynched/idempotent-transaction-api/domain/entities"
	"github.com/google/uuid"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) FindAllById(payee string) ([]entities.Transaction, error) {
	rows, err := r.db.Query("SELECT * FROM transactions WHERE payee_id = $1 OR payer_id = $1", payee)

	if err != nil {
		return nil, err
	}

	transactions := []entities.Transaction{}

	for rows.Next() {
		transaction := entities.Transaction{}

		err := rows.Scan(&transaction.Id, &transaction.Amount, &transaction.Payer, &transaction.Payee, &transaction.CreatedAt)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *TransactionRepository) Create(data entities.Transaction) (entities.Transaction, error) {
	tx, err := r.db.Begin()

	defer tx.Rollback()

	if err != nil {
		return entities.Transaction{}, err
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance - $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", data.Amount, data.Payer)

	if err != nil {
		return entities.Transaction{}, err
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance + $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", data.Amount, data.Payee)

	if err != nil {
		return entities.Transaction{}, err
	}

	id := uuid.NewString()

	row := tx.QueryRow("INSERT INTO transactions (id, amount, payer_id, payee_id) VALUES ($1, $2, $3, $4) RETURNING *", id, data.Amount, data.Payer, data.Payee)

	transaction := entities.Transaction{}

	err = row.Scan(&transaction.Id, &transaction.Amount, &transaction.Payer, &transaction.Payee, &transaction.CreatedAt)

	if err != nil {
		return entities.Transaction{}, err
	}

	if err := tx.Commit(); err != nil {
		return entities.Transaction{}, err
	}

	return transaction, nil
}
