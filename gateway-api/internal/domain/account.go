package domain

import (
	"github.com/lucsky/cuid"
	"github.com/swxtz/payment-gateway/gateway-api/internal/utils"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   string
	CreatedAt string
	UpdatedAt string
}

func NewAccount(name, email string) *Account {

	apiKey := utils.GenerateAPIKey()

	account := &Account{
		ID:     cuid.New(),
		Name:   name,
		Email:  email,
		APIKey: apiKey,
	}

	return account
}
