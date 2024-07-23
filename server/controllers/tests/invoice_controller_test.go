package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shota612/invoice-payment-service/server/controllers"
	"github.com/shota612/invoice-payment-service/server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockInvoiceUsecase struct {
	mock.Mock
}

func (m *MockInvoiceUsecase) CreateInvoice(issueDate string, paymentAmount, feeRate, salesTaxRate float64, paymentDueDate string, status models.InvoiceStatus, companyID, clientID uint) (*models.Invoice, error) {
	args := m.Called(issueDate, paymentAmount, feeRate, salesTaxRate, paymentDueDate, status, companyID, clientID)
	return args.Get(0).(*models.Invoice), args.Error(1)
}

func TestCreateInvoice(t *testing.T) {
	mockUsecase := new(MockInvoiceUsecase)
	controller := controllers.NewInvoiceController(mockUsecase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/invoices", controller.CreateInvoice)

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

	mockUsecase.On("CreateInvoice", invoice.IssueDate, invoice.PaymentAmount, invoice.FeeRate, invoice.SalesTaxRate, invoice.PaymentDueDate, invoice.Status, invoice.CompanyID, invoice.ClientID).Return(invoice, nil)

	w := httptest.NewRecorder()
	reqBody := `{
		"issue_date": "2024-07-23",
		"payment_amount": 10000,
		"fee_rate": 0.04,
		"sales_tax_rate": 0.10,
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

	var createdInvoice models.Invoice
	err := json.Unmarshal(w.Body.Bytes(), &createdInvoice)
	assert.Nil(t, err)
	assert.Equal(t, invoice.PaymentAmount, createdInvoice.PaymentAmount)
	assert.Equal(t, invoice.Fee, createdInvoice.Fee)
	assert.Equal(t, invoice.SalesTax, createdInvoice.SalesTax)
	assert.Equal(t, invoice.InvoiceAmount, createdInvoice.InvoiceAmount)
	assert.Equal(t, invoice.IssueDate, createdInvoice.IssueDate)
	assert.Equal(t, invoice.PaymentDueDate, createdInvoice.PaymentDueDate)
	assert.Equal(t, invoice.Status, createdInvoice.Status)
	assert.Equal(t, invoice.CompanyID, createdInvoice.CompanyID)
	assert.Equal(t, invoice.ClientID, createdInvoice.ClientID)
}
