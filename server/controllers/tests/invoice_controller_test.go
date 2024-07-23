package tests

import (
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

func (m *MockInvoiceUsecase) GetInvoicesByDateRange(startDate, endDate string) ([]models.Invoice, error) {
	args := m.Called(startDate, endDate)
	return args.Get(0).([]models.Invoice), args.Error(1)
}

func TestGetInvoicesByDateRange(t *testing.T) {
	mockUsecase := new(MockInvoiceUsecase)
	controller := controllers.NewInvoiceController(mockUsecase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/api/invoices", controller.GetInvoicesByDateRange)

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

	invoices := []models.Invoice{*invoice}

	mockUsecase.On("GetInvoicesByDateRange", "2024-07-01", "2024-07-31").Return(invoices, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/invoices?start_date=2024-07-01&end_date=2024-07-31", nil)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)

	var retrievedInvoices []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &retrievedInvoices)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(retrievedInvoices))

	expectedInvoice := map[string]interface{}{
		"ID":             invoice.ID,
		"IssueDate":      invoice.IssueDate,
		"PaymentAmount":  invoice.PaymentAmount,
		"Fee":            invoice.Fee,
		"FeeRate":        invoice.FeeRate,
		"SalesTax":       invoice.SalesTax,
		"SalesTaxRate":   invoice.SalesTaxRate,
		"InvoiceAmount":  invoice.InvoiceAmount,
		"PaymentDueDate": invoice.PaymentDueDate,
		"Status":         invoice.Status,
		"CompanyID":      float64(invoice.CompanyID),
		"ClientID":       float64(invoice.ClientID),
	}

	assert.Equal(t, expectedInvoice, retrievedInvoices[0])
}
