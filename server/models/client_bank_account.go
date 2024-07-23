package models

type ClientBankAccount struct {
	ID            uint   `gorm:"primaryKey"`
	ClientID      uint   `gorm:"not null"`
	BankName      string `gorm:"not null"`
	BranchName    string `gorm:"not null"`
	AccountNumber string `gorm:"not null"`
	AccountName   string `gorm:"not null"`
	Client        Client `gorm:"foreignKey:ClientID"`
}
