package usecase

import (
	"github.com/shota612/invoice-payment-service/server/domain"
	"github.com/shota612/invoice-payment-service/server/repository"
)

type InvoiceUsecase interface {
	CreateInvoice(issueDate string, paymentAmount float64, paymentDueDate string, status domain.InvoiceStatus, companyID, clientID uint) (*domain.Invoice, error)
	GetInvoicesByDateRange(startDate, endDate string) ([]domain.Invoice, error)
}

type invoiceUsecase struct {
	invoiceRepo repository.InvoiceRepository
}

func NewInvoiceUsecase(repo repository.InvoiceRepository) InvoiceUsecase {
	return &invoiceUsecase{repo}
}

func (u *invoiceUsecase) CreateInvoice(issueDate string, paymentAmount float64, paymentDueDate string, status domain.InvoiceStatus, companyID, clientID uint) (*domain.Invoice, error) {
	invoice := domain.NewInvoice(issueDate, paymentAmount, paymentDueDate, status, companyID, clientID)
	return u.invoiceRepo.CreateInvoice(invoice)
}

func (u *invoiceUsecase) GetInvoicesByDateRange(startDate, endDate string) ([]domain.Invoice, error) {
	return u.invoiceRepo.GetInvoicesByDateRange(startDate, endDate)
}
