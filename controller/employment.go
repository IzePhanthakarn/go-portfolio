package controller

import (
	"fmt"
	"net/http"

	"github.com/IzePhanthakarn/go-portfolio/db"
	"github.com/IzePhanthakarn/go-portfolio/dto"
	"github.com/IzePhanthakarn/go-portfolio/model"

	"github.com/gin-gonic/gin"
)

type Employment struct{}

func (p Employment) FindAll(ctx *gin.Context) {
	categoryId := ctx.Query("categoryId")
	search := ctx.Query("search")
	status := ctx.Query("status")

	var employments []model.Employment
	query := db.Conn.Preload("Category")
	if categoryId != "" {
		query = query.Where("category_id = ?", categoryId)
	}
	if search != "" {
		query = query.Where("sku = ? OR name LIKE ?", search, "%"+search+"%")

	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	query.Find(&employments)

	var result []dto.EmploymentResponse
	for _, employment := range employments {
		result = append(result, dto.EmploymentResponse{
			ID:          employment.ID,
			PlanName:    employment.PlanName,
			Darkmode:    employment.Darkmode,
			MultiLang:   employment.MultiLang,
			Name:        employment.Name,
			Email:       employment.Email,
			Phone:       employment.Phone,
			Price:       employment.Price,
			PriceTH:     employment.PriceTH,
			Status:      employment.Status,
			Description: employment.Description,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (p Employment) Create(ctx *gin.Context) {
	fmt.Println("0")
	var form dto.EmploymentRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(form)

	employment := model.Employment{
		PlanName:    form.PlanName,
		Darkmode:    form.Darkmode,
		MultiLang:   form.MultiLang,
		Name:        form.Name,
		Email:       form.Email,
		Phone:       form.Phone,
		Price:       form.Price,
		PriceTH:     form.PriceTH,
		Status:      0,
		Description: form.Description,
	}

	if err := db.Conn.Create(&employment).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.EmploymentResponse{
		ID:          employment.ID,
		PlanName:    employment.PlanName,
		Darkmode:    employment.Darkmode,
		MultiLang:   employment.MultiLang,
		Name:        employment.Name,
		Email:       employment.Email,
		Phone:       employment.Phone,
		Price:       employment.Price,
		PriceTH:     employment.PriceTH,
		Status:      employment.Status,
		Description: employment.Description,
	})
}
