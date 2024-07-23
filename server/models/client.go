package models

type Client struct {
	ID             uint                `gorm:"primaryKey"`
	LegalName      string              `gorm:"not null"`
	Representative string              `gorm:"not null"`
	PhoneNumber    string              `gorm:"not null"`
	PostalCode     string              `gorm:"not null"`
	Address        string              `gorm:"not null"`
	CompanyID      uint                `gorm:"not null"`
	Company        Company             `gorm:"foreignKey:CompanyID"`
	BankAccounts   []ClientBankAccount `gorm:"foreignKey:ClientID"`
	Invoices       []Invoice           `gorm:"foreignKey:ClientID"`
}
