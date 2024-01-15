package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model

	Name         string  `json:"name" gorm:"type:varchar(255);not null"`
	NameSocial   string  `json:"name_social" gorm:"type:varchar(255);not null"`
	Cnpj         string  `json:"cnpj" gorm:"type:varchar(14);not null;unique"`
	Contact      Contact `json:"contact" gorm:"embedded;embeddedPrefix:contact_"`
	MonthlyValue float64 `json:"monthly_value" gorm:"not null"`
	DueDate      int     `json:"due_date" gorm:"not null"`
	Address      Address `json:"address" gorm:"embedded;embeddedPrefix:address_"`
	Calc         *Calc   `json:"calc" gorm:"embedded;embeddedPrefix:calc_"`
}

type Contact struct {
	Email string `json:"email" gorm:"type:varchar(255);not null"`
	Phone string `json:"phone" gorm:"type:varchar(11);not null"`
}

type Address struct {
	Street       string `json:"street" gorm:"type:varchar(255);not null"`
	Number       string `json:"number" gorm:"type:varchar(255);not null"`
	Complement   string `json:"complement" gorm:"type:varchar(255)"`
	Neighborhood string `json:"neighborhood" gorm:"type:varchar(255);not null"`
	City         string `json:"city" gorm:"type:varchar(255);not null"`
	State        string `json:"state" gorm:"type:varchar(255);not null"`
	ZipCode      string `json:"zip_code" gorm:"type:varchar(8);not null"`
}
