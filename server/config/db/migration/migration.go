package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/moonrailgun/PUBGo/server/config/db"
	"github.com/moonrailgun/PUBGo/server/config/db/models"
)

// 执行迁移
func RunMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Player{}, &models.Match{}, &models.LifeTimeStats{}, &models.PubgRequestLog{})
}

func init() {
	RunMigrate(db.GetDb())
}
