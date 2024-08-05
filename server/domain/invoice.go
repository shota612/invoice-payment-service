package domain

type InvoiceStatus string

const (
	DefaultFeeRate      = 0.04
	DefaultSalesTaxRate = 0.10
)

const (
	Pending    InvoiceStatus = "Pending"
	Processing InvoiceStatus = "Processing"
	Paid       InvoiceStatus = "Paid"
	Error      InvoiceStatus = "Error"
)

type Invoice struct {
	ID             uint
	IssueDate      string
	PaymentAmount  float64
	Fee            float64
	FeeRate        float64
	SalesTax       float64
	SalesTaxRate   float64
	InvoiceAmount  float64
	PaymentDueDate string
	Status         InvoiceStatus
	CompanyID      uint
	ClientID       uint
}

func NewInvoice(issueDate string, paymentAmount float64, paymentDueDate string, status InvoiceStatus, companyID, clientID uint) *Invoice {
	invoice := &Invoice{
		IssueDate:      issueDate,
		PaymentAmount:  paymentAmount,
		FeeRate:        DefaultFeeRate,
		SalesTaxRate:   DefaultSalesTaxRate,
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
