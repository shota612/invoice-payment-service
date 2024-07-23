package tests

import (
	"github.com/shota612/invoice-payment-service/server/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientBankAccountModel(t *testing.T) {
	account := models.ClientBankAccount{
		ClientID:      1,
		BankName:      "Test Bank",
		BranchName:    "Test Branch",
		AccountNumber: "1234567890",
		AccountName:   "Test Account",
	}

	assert.Equal(t, uint(1), account.ClientID)
	assert.Equal(t, "Test Bank", account.BankName)
	assert.Equal(t, "Test Branch", account.BranchName)
	assert.Equal(t, "1234567890", account.AccountNumber)
	assert.Equal(t, "Test Account", account.AccountName)
}
