package gorm

import "gorm.io/gorm"

type ClientBankAccount struct {
	gorm.Model
	ClientID      uint
	Client        Client `gorm:"foreignKey:ClientID"`
	BankName      string `gorm:"type:varchar(255);not null"`
	BranchName    string `gorm:"type:varchar(255);not null"`
	AccountNumber string `gorm:"type:varchar(50);not null"`
	AccountName   string `gorm:"type:varchar(255);not null"`
}
