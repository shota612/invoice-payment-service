package repository

import (
	"github.com/shota612/invoice-payment-service/server/models"
	"gorm.io/gorm"
)

type InvoiceRepository interface {
	CreateInvoice(invoice *models.Invoice) (*models.Invoice, error)
	GetInvoicesByDateRange(startDate, endDate string) ([]models.Invoice, error)
}

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) InvoiceRepository {
	return &invoiceRepository{db}
}

func (r *invoiceRepository) CreateInvoice(invoice *models.Invoice) (*models.Invoice, error) {
	if err := r.db.Create(invoice).Error; err != nil {
		return nil, err
	}
	return invoice, nil
}

func (r *invoiceRepository) GetInvoicesByDateRange(startDate, endDate string) ([]models.Invoice, error) {
	var invoices []models.Invoice
	if err := r.db.Where("payment_due_date BETWEEN ? AND ?", startDate, endDate).Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}
