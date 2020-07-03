package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/moonrailgun/PUBGo/server/config"
	"github.com/moonrailgun/PUBGo/server/config/db/migration"
	"os"
)

// DB Global DB connection
var db *gorm.DB

func init() {
	var err error

	dbConfig := config.Config.DB
	db, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Name))
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	if os.Getenv("DEBUG") != "" {
		db.LogMode(true)
	}

	migration.RunMigrate(db) // 执行迁移
}

func GetDb() *gorm.DB {
	return db
}
