package models

import (
	"github.com/moonrailgun/PUBGo/server/config/db"
	"time"
)

type GlobalConfig struct {
	Name      string `gorm:"primary_key"`
	Value     string
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m *GlobalConfig) getConfigByKey(name string) error {
	return db.GetDb().Where("name = ?", name).First(m).Error
}

func (m *GlobalConfig) setConfig(name string, value string) error {
	m.Name = name
	m.Value = value
	return db.GetDb().Save(m).Error
}

// 获取全局设置
func GetGlobalConfig(name string) string {
	conf := new(GlobalConfig)
	err := conf.getConfigByKey(name)
	if err != nil {
		return ""
	}

	return conf.Value
}

// 设置全局设置
func SetGlobalConfig(name string, value string) error {
	conf := new(GlobalConfig)
	err := conf.setConfig(name, value)
	if err != nil {
		return err
	}

	return nil
}
