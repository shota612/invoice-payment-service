package tests

import (
	"github.com/shota612/invoice-payment-service/server/models"
	"github.com/shota612/invoice-payment-service/server/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	err := db.AutoMigrate(&models.Company{}, &models.User{}, &models.Client{}, &models.ClientBankAccount{}, &models.Invoice{})
	if err != nil {
		return nil
	}
	return db
}

func TestCreateInvoice(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewInvoiceRepository(db)

	invoice := models.NewInvoice(
		"2024-07-23",
		10000,
		"2024-08-23",
		models.Pending,
		1,
		1,
	)

	createdInvoice, err := repo.CreateInvoice(invoice)
	assert.Nil(t, err)
	assert.Equal(t, invoice.PaymentAmount, createdInvoice.PaymentAmount)
	assert.Equal(t, invoice.Fee, createdInvoice.Fee)
	assert.Equal(t, invoice.SalesTax, createdInvoice.SalesTax)
	assert.Equal(t, invoice.InvoiceAmount, createdInvoice.InvoiceAmount)

	var fetchedInvoice models.Invoice
	db.First(&fetchedInvoice, invoice.ID)

	assert.Equal(t, invoice.PaymentAmount, fetchedInvoice.PaymentAmount)
	assert.Equal(t, invoice.Fee, fetchedInvoice.Fee)
	assert.Equal(t, invoice.SalesTax, fetchedInvoice.SalesTax)
	assert.Equal(t, invoice.InvoiceAmount, fetchedInvoice.InvoiceAmount)
}

func TestGetInvoicesByDateRange(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewInvoiceRepository(db)

	invoice1 := models.NewInvoice(
		"2024-07-23",
		10000,
		"2024-07-31",
		models.Pending,
		1,
		1,
	)
	invoice2 := models.NewInvoice(
		"2024-08-01",
		20000,
		"2024-09-01",
		models.Pending,
		1,
		1,
	)

	_, err := repo.CreateInvoice(invoice1)
	if err != nil {
		return
	}
	_, err = repo.CreateInvoice(invoice2)
	if err != nil {
		return
	}

	startDate := "2024-07-01"
	endDate := "2024-07-31"

	invoices, err := repo.GetInvoicesByDateRange(startDate, endDate)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(invoices))
	assert.Equal(t, invoice1.IssueDate, invoices[0].IssueDate)
}
