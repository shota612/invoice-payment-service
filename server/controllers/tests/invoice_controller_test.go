package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shota612/invoice-payment-service/server/controllers"
	"github.com/shota612/invoice-payment-service/server/controllers/adapter"
	"github.com/shota612/invoice-payment-service/server/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockInvoiceUsecase struct {
	mock.Mock
}

func (m *MockInvoiceUsecase) CreateInvoice(issueDate string, paymentAmount float64, paymentDueDate string, status domain.InvoiceStatus, companyID, clientID uint) (*domain.Invoice, error) {
	args := m.Called(issueDate, paymentAmount, paymentDueDate, status, companyID, clientID)
	return args.Get(0).(*domain.Invoice), args.Error(1)
}

func (m *MockInvoiceUsecase) GetInvoicesByDateRange(startDate, endDate string) ([]domain.Invoice, error) {
	args := m.Called(startDate, endDate)
	return args.Get(0).([]domain.Invoice), args.Error(1)
}

func TestCreateInvoice(t *testing.T) {
	mockUsecase := new(MockInvoiceUsecase)
	controller := controllers.NewInvoiceController(mockUsecase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/invoices", controller.CreateInvoice)

	invoice := domain.NewInvoice(
		"2024-07-23",
		10000,
		"2024-08-23",
		domain.Pending,
		1,
		1,
	)

	mockUsecase.On("CreateInvoice", invoice.IssueDate, invoice.PaymentAmount, invoice.PaymentDueDate, invoice.Status, invoice.CompanyID, invoice.ClientID).Return(invoice, nil)

	w := httptest.NewRecorder()
	reqBody := `{
		"issue_date": "2024-07-23",
		"payment_amount": 10000,
		"payment_due_date": "2024-08-23",
		"status": "Pending",
		"company_id": 1,
		"client_id": 1
	}`
	req, _ := http.NewRequest("POST", "/api/invoices", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)

	var createdInvoice adapter.InvoiceResponse
	err := json.Unmarshal(w.Body.Bytes(), &createdInvoice)
	assert.Nil(t, err)
	assert.Equal(t, invoice.ID, createdInvoice.ID)
	assert.Equal(t, invoice.IssueDate, createdInvoice.IssueDate)
	assert.Equal(t, invoice.PaymentAmount, createdInvoice.PaymentAmount)
	assert.Equal(t, invoice.Fee, createdInvoice.Fee)
	assert.Equal(t, invoice.SalesTax, createdInvoice.SalesTax)
	assert.Equal(t, invoice.InvoiceAmount, createdInvoice.InvoiceAmount)
	assert.Equal(t, invoice.PaymentDueDate, createdInvoice.PaymentDueDate)
	assert.Equal(t, invoice.Status, createdInvoice.Status)
	assert.Equal(t, invoice.CompanyID, createdInvoice.CompanyID)
	assert.Equal(t, invoice.ClientID, createdInvoice.ClientID)
}

func TestGetInvoicesByDateRange(t *testing.T) {
	mockUsecase := new(MockInvoiceUsecase)
	controller := controllers.NewInvoiceController(mockUsecase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/api/invoices", controller.GetInvoicesByDateRange)

	invoice := domain.NewInvoice(
		"2024-07-23",
		10000,
		"2024-08-23",
		domain.Pending,
		1,
		1,
	)

	invoices := []domain.Invoice{*invoice}

	mockUsecase.On("GetInvoicesByDateRange", "2024-07-01", "2024-07-31").Return(invoices, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/invoices?start_date=2024-07-01&end_date=2024-07-31", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)

	var retrievedInvoices []adapter.InvoiceResponse
	err := json.Unmarshal(w.Body.Bytes(), &retrievedInvoices)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(retrievedInvoices))

	expectedInvoice := adapter.InvoiceResponse{
		ID:             invoice.ID,
		IssueDate:      invoice.IssueDate,
		PaymentAmount:  invoice.PaymentAmount,
		Fee:            invoice.Fee,
		FeeRate:        invoice.FeeRate,
		SalesTax:       invoice.SalesTax,
		SalesTaxRate:   invoice.SalesTaxRate,
		InvoiceAmount:  invoice.InvoiceAmount,
		PaymentDueDate: invoice.PaymentDueDate,
		Status:         invoice.Status,
		CompanyID:      invoice.CompanyID,
		ClientID:       invoice.ClientID,
	}

	assert.Equal(t, expectedInvoice, retrievedInvoices[0])
}

// TODO: Add abnormally test cases like error cases
