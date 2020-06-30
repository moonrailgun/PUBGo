package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/moonrailgun/PUBGo/server/config"
	"os"
)

// DB Global DB connection
var DB *gorm.DB

func init() {
	var err error

	dbConfig := config.Config.DB
	DB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)); if err != nil {
		panic(err)
	}

	if os.Getenv("DEBUG") != "" {
		DB.LogMode(true)
	}
}