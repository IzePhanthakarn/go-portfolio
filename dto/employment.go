package dto

type EmploymentRequest struct {
	PlanName    string  `form:"planName" binding:"required"`
	Darkmode    uint    `form:"darkmode"`
	MultiLang   uint    `form:"multiLang"`
	Name        string  `form:"name" binding:"required"`
	Email       string  `form:"email" binding:"required"`
	Phone       string  `form:"phone" binding:"required"`
	Price       float64 `form:"price" binding:"required"`
	PriceTH     uint    `form:"priceTH" binding:"required"`
	Status      uint    `form:"status"`
	Description string  `form:"description"`
}

type EmploymentResponse struct {
	ID          uint    `json:"id"`
	PlanName    string  `json:"planName"`
	Darkmode    uint    `json:"darkmode"`
	MultiLang   uint    `json:"multiLang"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Price       float64 `json:"price"`
	PriceTH     uint    `json:"priceTH"`
	Status      uint    `json:"status"`
	Description string  `json:"description"`
}
