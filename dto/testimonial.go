package dto

import (
	"time"
)

type TestimonialRequest struct {
	Name         string `form:"name" binding:"required"`
	Relationship string `form:"relationship" binding:"required"`
	// Image        *multipart.FileHeader `form:"image" binding:"required"`
	Description string  `form:"description" binding:"required"`
	Score       float64 `form:"score" binding:"required"`
}

type CreateOrUpdateTestimonialResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Relationship string  `json:"relationship"`
	Image        string  `json:"image"`
	Description  string  `json:"description"`
	Score        float64 `json:"score"`
	Status       uint    `json:"status"`
}

type ReadTestimonialResponse struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	Name         string    `json:"name"`
	Relationship string    `json:"relationship"`
	Image        string    `json:"image"`
	Description  string    `json:"description"`
	Score        float64   `json:"score"`
	Status       uint      `json:"status"`
}
