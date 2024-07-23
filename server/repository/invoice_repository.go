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
	err := r.db.Create(invoice).Error
	return invoice, err
}

func (r *invoiceRepository) GetInvoicesByDateRange(startDate, endDate string) ([]models.Invoice, error) {
	var invoices []models.Invoice
	err := r.db.Where("payment_due_date BETWEEN ? AND ?", startDate, endDate).Find(&invoices).Error
	return invoices, err
}
