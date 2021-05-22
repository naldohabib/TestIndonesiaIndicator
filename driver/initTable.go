package driver

import (
	"TestScrapeCRUD/models"
	"github.com/jinzhu/gorm"
)

// InitTable ...
func InitTable(db *gorm.DB) {
	db.Debug().AutoMigrate(
		&models.YoutubeData{},
	)
}
