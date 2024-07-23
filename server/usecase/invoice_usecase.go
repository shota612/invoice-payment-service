package usecase

import (
	"github.com/shota612/invoice-payment-service/server/models"
	"github.com/shota612/invoice-payment-service/server/repository"
)

type InvoiceUsecase interface {
	CreateInvoice(issueDate string, paymentAmount float64, paymentDueDate string, status models.InvoiceStatus, companyID, clientID uint) (*models.Invoice, error)
	GetInvoicesByDateRange(startDate, endDate string) ([]models.Invoice, error)
}

type invoiceUsecase struct {
	invoiceRepo repository.InvoiceRepository
}

func NewInvoiceUsecase(repo repository.InvoiceRepository) InvoiceUsecase {
	return &invoiceUsecase{repo}
}

func (u *invoiceUsecase) CreateInvoice(issueDate string, paymentAmount float64, paymentDueDate string, status models.InvoiceStatus, companyID, clientID uint) (*models.Invoice, error) {
	invoice := models.NewInvoice(issueDate, paymentAmount, paymentDueDate, status, companyID, clientID)
	return u.invoiceRepo.CreateInvoice(invoice)
}

func (u *invoiceUsecase) GetInvoicesByDateRange(startDate, endDate string) ([]models.Invoice, error) {
	return u.invoiceRepo.GetInvoicesByDateRange(startDate, endDate)
}
