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

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateEmploymentResponse{
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
