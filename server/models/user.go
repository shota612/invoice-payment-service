package models

type User struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"not null"`
	Email     string  `gorm:"not null;unique"`
	Password  string  `gorm:"not null"`
	CompanyID uint    `gorm:"not null"`
	Company   Company `gorm:"foreignKey:CompanyID"`
}
