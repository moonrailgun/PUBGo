package models

import (
	"github.com/jinzhu/gorm"
	"github.com/moonrailgun/PUBGo/server/config/db"
	"github.com/moonrailgun/PUBGo/server/config/pubg"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
	"github.com/moonrailgun/PUBGo/server/config/pubg/schema"
	"time"
)

type ModelPlayer struct {
	gorm.Model
	AccountId        string `gorm:"not null;unique_index:index_account"`
	Name             string `gorm:"not null;unique_index:index_account"`
	ShardID          string
	AccountCreatedAt time.Time
	AccountUpdatedAt time.Time
}

func (m *ModelPlayer) ParseFromPUBG(data schema.Player) {
	m.AccountId = data.ID
	m.Name = data.Name
	m.ShardID = data.ShardID
	m.AccountCreatedAt = data.CreatedAt
	m.AccountUpdatedAt = data.UpdatedAt
}

// 根据用户名获取用户信息
// 优先从数据库中获取
func (m *ModelPlayer) GetInfoByUserName(shard api.ShardType, username string) error {
	err := db.GetDb().Where("shard_id = ? and name = ?", shard, username).First(m)
	if err != nil {
		// 	数据库没有数据，发请求获取
		remotePlayer, err := pubg.Api.RequestSinglePlayerByName(shard, username)
		if err != nil {
			return err
		}
		m.ParseFromPUBG(*remotePlayer)
		db.GetDb().Create(m)
	}

	return nil
}
