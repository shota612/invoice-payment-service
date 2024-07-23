package tests

import (
	"github.com/shota612/invoice-payment-service/server/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompanyModel(t *testing.T) {
	company := models.Company{
		LegalName:      "Test Company",
		Representative: "John Doe",
		PhoneNumber:    "123456789",
		PostalCode:     "123-4567",
		Address:        "123 Test St, Test City",
	}

	assert.Equal(t, "Test Company", company.LegalName)
	assert.Equal(t, "John Doe", company.Representative)
	assert.Equal(t, "123456789", company.PhoneNumber)
	assert.Equal(t, "123-4567", company.PostalCode)
	assert.Equal(t, "123 Test St, Test City", company.Address)
}
