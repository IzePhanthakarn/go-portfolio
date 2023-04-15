package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/IzePhanthakarn/go-portfolio/db"
	"github.com/IzePhanthakarn/go-portfolio/dto"
	"github.com/IzePhanthakarn/go-portfolio/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Employment struct{}

func (p Employment) FindAll(ctx *gin.Context) {
	name := ctx.Query("name")
	email := ctx.Query("email")
	plan := ctx.Query("planName")
	status := ctx.Query("status")

	var employments []model.Employment
	query := db.Conn
	if name != "" {
		query = query.Where("name = ? OR name LIKE ?", name, "%"+name+"%")
	}
	if email != "" {
		query = query.Where("email = ? OR email LIKE ?", email, "%"+email+"%")

	}
	if plan != "" {
		query = query.Where("plan_name = ?", plan)
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

func (p Employment) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var employment model.Employment

	query := db.Conn.First(&employment, id)
	if err := query.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.EmploymentResponse{
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

func (p Employment) Create(ctx *gin.Context) {
	var form dto.EmploymentRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
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

func (p Employment) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var form dto.EmploymentRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var employment model.Employment
	if err := db.Conn.First(&employment, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	employment.PlanName = form.PlanName
	employment.Darkmode = form.Darkmode
	employment.MultiLang = form.MultiLang
	employment.Name = form.Name
	employment.Email = form.Email
	employment.Phone = form.Phone
	employment.Price = form.Price
	employment.PriceTH = form.PriceTH
	employment.Status = form.Status
	employment.Description = form.Description
	db.Conn.Save(&employment)

	ctx.JSON(http.StatusOK, dto.EmploymentResponse{
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

func (p Employment) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	db.Conn.Unscoped().Delete(&model.Employment{}, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
