package repository

import (
	"database/sql"
	"time"

	"github.com/swxtz/payment-gateway/gateway-api/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) Save(account *domain.Account) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(account.ID, account.Name, account.Email, account.APIKey, account.Balance, account.CreatedAt, account.UpdatedAt)

	if err != nil {
		panic(err)
	}

	return nil
}

func (r *AccountRepository) FindByApiKey(apiKey string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(
		`SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE api_key = ?`,
		apiKey).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&createdAt,
		&updatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt

	return &account, nil
}

func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(`
		SELECT * FROM accounts WHERE id = ?`,
	).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&createdAt,
		&updatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt
	return &account, nil
}
