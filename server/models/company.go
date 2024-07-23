package models

type Company struct {
	ID             uint      `gorm:"primaryKey"`
	LegalName      string    `gorm:"not null"`
	Representative string    `gorm:"not null"`
	PhoneNumber    string    `gorm:"not null"`
	PostalCode     string    `gorm:"not null"`
	Address        string    `gorm:"not null"`
	Users          []User    `gorm:"foreignKey:CompanyID"`
	Clients        []Client  `gorm:"foreignKey:CompanyID"`
	Invoices       []Invoice `gorm:"foreignKey:CompanyID"`
}
