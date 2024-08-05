package gorm

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model
	IssueDate      string  `gorm:"type:date;not null"`
	PaymentAmount  float64 `gorm:"type:real;not null"`
	Fee            float64 `gorm:"type:real;not null"`
	FeeRate        float64 `gorm:"type:real;not null"`
	SalesTax       float64 `gorm:"type:real;not null"`
	SalesTaxRate   float64 `gorm:"type:real;not null"`
	InvoiceAmount  float64 `gorm:"type:real;not null"`
	PaymentDueDate string  `gorm:"type:date;not null"`
	Status         string  `gorm:"type:varchar(50);not null;check:status IN ('Pending', 'Processing', 'Paid', 'Error')"`
	CompanyID      uint
	Company        Company `gorm:"foreignKey:CompanyID"`
	ClientID       uint
	Client         Client `gorm:"foreignKey:ClientID"`
}
