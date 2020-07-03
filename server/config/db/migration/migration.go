package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/moonrailgun/PUBGo/server/config/db/models"
)

// 执行迁移
func RunMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.ModelPlayer{}, &models.ModelMatch{})
}
