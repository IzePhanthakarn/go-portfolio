package controller

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/IzePhanthakarn/go-portfolio/db"
	"github.com/IzePhanthakarn/go-portfolio/dto"
	"github.com/IzePhanthakarn/go-portfolio/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Testimonial struct{}

func (t Testimonial) FindAll(ctx *gin.Context) {
	name := ctx.Query("name")
	relationship := ctx.Query("relationship")
	createdAt := ctx.Query("createdAt")
	score := ctx.Query("score")

	var testimonials []model.Testimonial
	query := db.Conn
	if name != "" {
		query = query.Where("name = ? OR name LIKE ?", name, "%"+name+"%")
	}
	if relationship != "" {
		query = query.Where("relationship = ? OR relationship LIKE ?", relationship, "%"+relationship+"%")
	}
	if createdAt != "" {
		query = query.Where("created_at = ?", createdAt)
	}
	if score != "" {
		query = query.Where("score = ?", score)
	}
	query.Find(&testimonials)

	var result []dto.ReadTestimonialResponse
	for _, testimonial := range testimonials {
		result = append(result, dto.ReadTestimonialResponse{
			ID:           testimonial.ID,
			CreatedAt:    testimonial.CreatedAt,
			Name:         testimonial.Name,
			Relationship: testimonial.Relationship,
			Image:        testimonial.Image,
			Description:  testimonial.Description,
			Score:        testimonial.Score,
			Status:       testimonial.Status,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (t Testimonial) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var testimonial model.Testimonial

	query := db.Conn.First(&testimonial, id)
	if err := query.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ReadTestimonialResponse{
		ID:           testimonial.ID,
		CreatedAt:    testimonial.CreatedAt,
		Name:         testimonial.Name,
		Relationship: testimonial.Relationship,
		Image:        testimonial.Image,
		Description:  testimonial.Description,
		Score:        testimonial.Score,
		Status:       testimonial.Status,
	})
}

func (t Testimonial) Create(ctx *gin.Context) {
	var form dto.TestimonialRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imagePath := "./uploads/testimonials/" + uuid.New().String()
	ctx.SaveUploadedFile(image, imagePath)

	testimonial := model.Testimonial{
		Name:         form.Name,
		Relationship: form.Relationship,
		Description:  form.Description,
		Score:        form.Score,
		Status:       0,
		Image:        imagePath,
	}

	if err := db.Conn.Create(&testimonial).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateOrUpdateTestimonialResponse{
		ID:           testimonial.ID,
		Name:         testimonial.Name,
		Relationship: testimonial.Relationship,
		Image:        testimonial.Image,
		Description:  testimonial.Description,
		Score:        testimonial.Score,
		Status:       testimonial.Status,
	})
}

func (p Testimonial) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var form dto.TestimonialRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var testimonial model.Testimonial
	if err := db.Conn.First(&testimonial, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if image != nil {
		imagePath := "./uploads/testimonials/" + uuid.New().String()
		ctx.SaveUploadedFile(image, imagePath)
		os.Remove(testimonial.Image)
		testimonial.Image = imagePath
	}
	testimonial.Name = form.Name
	db.Conn.Save(&testimonial)

	ctx.JSON(http.StatusOK, dto.CreateOrUpdateTestimonialResponse{
		ID:           testimonial.ID,
		Name:         testimonial.Name,
		Relationship: testimonial.Relationship,
		Image:        testimonial.Image,
		Description:  testimonial.Description,
		Score:        testimonial.Score,
		Status:       testimonial.Status,
	})
}

func (t Testimonial) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	db.Conn.Unscoped().Delete(&model.Testimonial{}, id)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
