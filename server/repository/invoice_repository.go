package repository

import (
	"github.com/shota612/invoice-payment-service/server/domain"
	"gorm.io/gorm"
)

type InvoiceRepository interface {
	CreateInvoice(invoice *domain.Invoice) (*domain.Invoice, error)
	GetInvoicesByDateRange(startDate, endDate string) ([]domain.Invoice, error)
}

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) InvoiceRepository {
	return &invoiceRepository{db}
}

func (r *invoiceRepository) CreateInvoice(invoice *domain.Invoice) (*domain.Invoice, error) {
	if err := r.db.Create(invoice).Error; err != nil {
		return nil, err
	}
	return invoice, nil
}

func (r *invoiceRepository) GetInvoicesByDateRange(startDate, endDate string) ([]domain.Invoice, error) {
	var invoices []domain.Invoice
	if err := r.db.Where("payment_due_date BETWEEN ? AND ?", startDate, endDate).Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}
