package tests

import (
	"testing"

	"github.com/shota612/invoice-payment-service/server/models"
	"github.com/shota612/invoice-payment-service/server/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockInvoiceRepository struct {
	mock.Mock
}

func (m *MockInvoiceRepository) CreateInvoice(invoice *models.Invoice) (*models.Invoice, error) {
	args := m.Called(invoice)
	return args.Get(0).(*models.Invoice), args.Error(1)
}

func (m *MockInvoiceRepository) GetInvoicesByDateRange(startDate, endDate string) ([]models.Invoice, error) {
	args := m.Called(startDate, endDate)
	return args.Get(0).([]models.Invoice), args.Error(1)
}

func TestCreateInvoice(t *testing.T) {
	mockRepo := new(MockInvoiceRepository)
	invoiceUsecase := usecase.NewInvoiceUsecase(mockRepo)

	invoice := models.NewInvoice(
		"2024-07-23",
		10000,
		"2024-08-23",
		models.Pending,
		1,
		1,
	)

	mockRepo.On("CreateInvoice", invoice).Return(invoice, nil)

	createdInvoice, err := invoiceUsecase.CreateInvoice(
		"2024-07-23",
		10000,
		"2024-08-23",
		models.Pending,
		1,
		1,
	)

	assert.Nil(t, err)
	assert.Equal(t, invoice, createdInvoice)
	mockRepo.AssertExpectations(t)
}

func TestGetInvoicesByDateRange(t *testing.T) {
	mockRepo := new(MockInvoiceRepository)
	invoiceUsecase := usecase.NewInvoiceUsecase(mockRepo)

	invoice1 := models.NewInvoice(
		"2024-07-23",
		10000,
		"2024-08-23",
		models.Pending,
		1,
		1,
	)

	invoices := []models.Invoice{*invoice1}

	mockRepo.On("GetInvoicesByDateRange", "2024-07-01", "2024-07-31").Return(invoices, nil)

	startDate := "2024-07-01"
	endDate := "2024-07-31"

	retrievedInvoices, err := invoiceUsecase.GetInvoicesByDateRange(startDate, endDate)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(retrievedInvoices))
	assert.Equal(t, invoice1.ID, retrievedInvoices[0].ID)
	mockRepo.AssertExpectations(t)
}
