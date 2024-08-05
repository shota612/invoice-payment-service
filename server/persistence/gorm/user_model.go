package gorm

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	CompanyID uint
	Company   Company `gorm:"foreignKey:CompanyID"`
}
