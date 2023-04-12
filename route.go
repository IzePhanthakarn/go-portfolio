package main

import (
	"github.com/IzePhanthakarn/go-portfolio/controller"

	"github.com/gin-gonic/gin"
)

func serveRoutes(r *gin.Engine) {
	employmentController := controller.Employment{}
	employmentGroup := r.Group("/employments")
	employmentGroup.POST("", employmentController.Create)
}
