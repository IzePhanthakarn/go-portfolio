package model

import "gorm.io/gorm"

type Employment struct {
	gorm.Model
	PlanName string `gorm:"type:varchar(100);not null"`
	Darkmode uint `gorm:"not null"`
	MultiLang uint `gorm:"not null"`
	Name string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);not null"`
	Phone string `gorm:"type:varchar(10);not null"`
	Price float64 `gorm:"not null"`
	PriceTH uint `gorm:"not null"`
	Status uint `gorm:"not null"`
	Description string `gorm:"type:varchar(1000);not null"`
}
