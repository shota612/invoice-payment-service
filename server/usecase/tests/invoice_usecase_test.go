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

func TestCreateInvoice(t *testing.T) {
	mockRepo := new(MockInvoiceRepository)
	invoiceUsecase := usecase.NewInvoiceUsecase(mockRepo)

	invoice := models.NewInvoice(
		"2024-07-23",
		10000,
		0.04,
		0.10,
		"2024-08-23",
		models.Pending,
		1,
		1,
	)

	mockRepo.On("CreateInvoice", invoice).Return(invoice, nil)

	createdInvoice, err := invoiceUsecase.CreateInvoice(
		"2024-07-23",
		10000,
		0.04,
		0.10,
		"2024-08-23",
		models.Pending,
		1,
		1,
	)

	assert.Nil(t, err)
	assert.Equal(t, invoice, createdInvoice)
	mockRepo.AssertExpectations(t)
}
