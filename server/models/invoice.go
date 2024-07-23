package models

type InvoiceStatus string

const (
	Pending    InvoiceStatus = "Pending"
	Processing InvoiceStatus = "Processing"
	Paid       InvoiceStatus = "Paid"
	Error      InvoiceStatus = "Error"
)

type Invoice struct {
	ID             uint          `gorm:"primaryKey"`
	IssueDate      string        `gorm:"not null"`
	PaymentAmount  float64       `gorm:"not null"`
	Fee            float64       `gorm:"-"`
	FeeRate        float64       `gorm:"not null"`
	SalesTax       float64       `gorm:"-"`
	SalesTaxRate   float64       `gorm:"not null"`
	InvoiceAmount  float64       `gorm:"-"`
	PaymentDueDate string        `gorm:"not null"`
	Status         InvoiceStatus `gorm:"not null"`
	CompanyID      uint          `gorm:"not null"`
	Company        Company       `gorm:"foreignKey:CompanyID"`
	ClientID       uint          `gorm:"not null"`
	Client         Client        `gorm:"foreignKey:ClientID"`
}

func NewInvoice(issueDate string, paymentAmount, feeRate, salesTaxRate float64, paymentDueDate string, status InvoiceStatus, companyID, clientID uint) *Invoice {
	invoice := &Invoice{
		IssueDate:      issueDate,
		PaymentAmount:  paymentAmount,
		FeeRate:        feeRate,
		SalesTaxRate:   salesTaxRate,
		PaymentDueDate: paymentDueDate,
		Status:         status,
		CompanyID:      companyID,
		ClientID:       clientID,
	}
	invoice.CalculateAmounts()
	return invoice
}

func (i *Invoice) CalculateAmounts() {
	i.Fee = i.PaymentAmount * i.FeeRate
	i.SalesTax = i.Fee * i.SalesTaxRate
	i.InvoiceAmount = i.PaymentAmount + i.Fee + i.SalesTax
}
