package domain

import (
	"sync"
	"time"

	"github.com/lucsky/cuid"
	"github.com/swxtz/payment-gateway/gateway-api/internal/utils"
)

type Account struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	APIKey    string  `json:"apiKey"`
	Balance   float64 `json:"balance"`
	mu        sync.RWMutex
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(name, email string) *Account {

	apiKey := utils.GenerateAPIKey()

	account := &Account{
		ID:        cuid.New(),
		Name:      name,
		Email:     email,
		APIKey:    apiKey,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func (a *Account) AddBalance(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Balance += amount
	a.UpdatedAt = time.Now()
}
