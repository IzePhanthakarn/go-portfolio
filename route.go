package main

import (
	"github.com/IzePhanthakarn/go-portfolio/controller"

	"github.com/gin-gonic/gin"
)

func serveRoutes(r *gin.Engine) {
	employmentController := controller.Employment{}
	employmentGroup := r.Group("/employments")
	employmentGroup.GET("", employmentController.FindAll)
	employmentGroup.GET("/:id", employmentController.FindOne)
	employmentGroup.POST("", employmentController.Create)
	employmentGroup.PATCH("/:id", employmentController.Update)
	employmentGroup.DELETE("/:id", employmentController.Delete)

	testimonialController := controller.Testimonial{}
	testimonialGroup := r.Group("/testimonials")
	testimonialGroup.GET("", testimonialController.FindAll)
	testimonialGroup.GET("/:id", testimonialController.FindOne)
	testimonialGroup.POST("", testimonialController.Create)
	testimonialGroup.PATCH("/:id", testimonialController.Update)
	testimonialGroup.DELETE("/:id", testimonialController.Delete)
}