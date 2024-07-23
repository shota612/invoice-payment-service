package tests

import (
	"github.com/shota612/invoice-payment-service/server/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvoiceModel(t *testing.T) {
	invoice := models.NewInvoice(
		"2024-07-23",
		10000,
		"2024-08-23",
		models.Pending,
		1,
		1,
	)

	assert.Equal(t, 10000.0, invoice.PaymentAmount)
	assert.Equal(t, 400.0, invoice.Fee)
	assert.Equal(t, 40.0, invoice.SalesTax)
	assert.Equal(t, 10440.0, invoice.InvoiceAmount)
	assert.Equal(t, "2024-07-23", invoice.IssueDate)
	assert.Equal(t, "2024-08-23", invoice.PaymentDueDate)
	assert.Equal(t, models.Pending, invoice.Status)
	assert.Equal(t, uint(1), invoice.CompanyID)
	assert.Equal(t, uint(1), invoice.ClientID)
}
