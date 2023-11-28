package initializers

import (
	"github.com/robbyklein/trivia-backend/models"
	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Question{})
}
