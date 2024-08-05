package gorm

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	LegalName          string `gorm:"type:varchar(255);not null"`
	RepresentativeName string `gorm:"type:varchar(255);not null"`
	PhoneNumber        string `gorm:"type:varchar(20);not null"`
	PostalCode         string `gorm:"type:varchar(10);not null"`
	Address            string `gorm:"type:text;not null"`
}
