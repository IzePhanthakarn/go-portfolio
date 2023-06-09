package db

import (
	"log"
	"os"

	"github.com/IzePhanthakarn/go-portfolio/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	if err != nil {
		log.Fatal("Cannot connect to the database")
	}

	Conn = db
}

func Migrate() {
	Conn.AutoMigrate(
		&model.Employment{},
		&model.Testimonial{},
	)
}
