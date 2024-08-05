package adapter

import "github.com/shota612/invoice-payment-service/server/domain"

type InvoiceResponse struct {
	ID             uint                 `json:"id"`
	IssueDate      string               `json:"issue_date"`
	PaymentAmount  float64              `json:"payment_amount"`
	Fee            float64              `json:"fee"`
	FeeRate        float64              `json:"fee_rate"`
	SalesTax       float64              `json:"sales_tax"`
	SalesTaxRate   float64              `json:"sales_tax_rate"`
	InvoiceAmount  float64              `json:"invoice_amount"`
	PaymentDueDate string               `json:"payment_due_date"`
	Status         domain.InvoiceStatus `json:"status"`
	CompanyID      uint                 `json:"company_id"`
	ClientID       uint                 `json:"client_id"`
}

func NewInvoiceResponse(invoice domain.Invoice) InvoiceResponse {
	return InvoiceResponse{
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
}

func NewInvoiceResponses(invoices []domain.Invoice) []InvoiceResponse {
	responses := make([]InvoiceResponse, len(invoices))
	for i, invoice := range invoices {
		responses[i] = NewInvoiceResponse(invoice)
	}
	return responses
}
