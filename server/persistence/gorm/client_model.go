package gorm

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	LegalName          string `gorm:"type:varchar(255);not null"`
	RepresentativeName string `gorm:"type:varchar(255);not null"`
	PhoneNumber        string `gorm:"type:varchar(20);not null"`
	PostalCode         string `gorm:"type:varchar(10);not null"`
	Address            string `gorm:"type:text;not null"`
	CompanyID          uint
	Company            Company             `gorm:"foreignKey:CompanyID"`
	ClientBankAccounts []ClientBankAccount `gorm:"foreignKey:ClientID"`
	Invoices           []Invoice           `gorm:"foreignKey:ClientID"`
}
