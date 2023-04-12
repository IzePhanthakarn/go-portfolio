package main

import (
	"github.com/IzePhanthakarn/go-portfolio/db"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	db.ConnectDB()
	db.Migrate()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	os.MkdirAll("uploads/testimonials", 0755)
	r := gin.Default()
	r.Use(cors.New(corsConfig))
	r.Static("/testimonials", "./uploads")
	serveRoutes(r)
	r.Run(":" + os.Getenv("PORT"))
}