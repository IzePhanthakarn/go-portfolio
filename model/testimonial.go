package model

import "gorm.io/gorm"

type Testimonial struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(100);not null"`
	Relationship string  `gorm:"type:varchar(100);not null"`
	Image        string  `gorm:"type:varchar(255);not null"`
	Description  string  `gorm:"type:varchar(1000);not null"`
	Score        float64 `gorm:"not null"`
	Status       uint    `gorm:"not null"`
}
