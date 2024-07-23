package tests

import (
	"github.com/shota612/invoice-payment-service/server/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientModel(t *testing.T) {
	client := models.Client{
		LegalName:      "Client Company",
		Representative: "Jane Doe",
		PhoneNumber:    "987654321",
		PostalCode:     "765-4321",
		Address:        "321 Test Ave, Test City",
		CompanyID:      1,
	}

	assert.Equal(t, "Client Company", client.LegalName)
	assert.Equal(t, "Jane Doe", client.Representative)
	assert.Equal(t, "987654321", client.PhoneNumber)
	assert.Equal(t, "765-4321", client.PostalCode)
	assert.Equal(t, "321 Test Ave, Test City", client.Address)
	assert.Equal(t, uint(1), client.CompanyID)
}
