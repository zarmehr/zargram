package db

import (
	"log"
	"zargram/models"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Friend{}, &models.Story{}, &models.Reaction{})
	if err != nil {
		log.Fatal(err)
	}
}
