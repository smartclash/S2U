package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/smartclash/S2U/database/models"
)

var Database *gorm.DB

func RunMigrations() {
	Database.AutoMigrate(&models.User{})
}
