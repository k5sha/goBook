package storage

import (
	"github.com/k5sha/goBook/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
