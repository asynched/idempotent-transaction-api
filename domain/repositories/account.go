package repositories

import (
	"database/sql"

	"github.com/asynched/idempotent-transaction-api/domain/entities"
	"github.com/google/uuid"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (r *AccountRepository) Save(data entities.Account) (entities.Account, error) {
	id := uuid.NewString()

	row := r.db.QueryRow("INSERT INTO accounts (id, name, cpf) VALUES ($1, $2, $3) RETURNING *", id, data.Name, data.Cpf)

	account := entities.Account{}

	err := row.Scan(&account.Id, &account.Name, &account.Cpf, &account.Balance, &account.CreatedAt, &account.UpdatedAt)

	if err != nil {
		return entities.Account{}, err
	}

	return account, nil
}

func (r *AccountRepository) FindById(id string) (entities.Account, error) {
	row := r.db.QueryRow("SELECT * FROM accounts WHERE id = $1", id)

	account := entities.Account{}

	err := row.Scan(&account.Id, &account.Name, &account.Cpf, &account.Balance, &account.CreatedAt, &account.UpdatedAt)

	if err != nil {
		return entities.Account{}, err
	}

	return account, nil
}
