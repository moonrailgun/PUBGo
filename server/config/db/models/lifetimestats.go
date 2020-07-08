package models

import (
	"github.com/jinzhu/gorm"
	"github.com/moonrailgun/PUBGo/server/config/db"
	"github.com/moonrailgun/PUBGo/server/config/pubg"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
	"github.com/moonrailgun/PUBGo/server/config/pubg/schema"
	"github.com/moonrailgun/PUBGo/server/utils"
)

type ModelLifeTimeStats struct {
	gorm.Model
	AccountId       string
	GameModeStats   string `gorm:"type:blob"`
	MatchesSolo     string `gorm:"type:blob"`
	MatchesSoloFPP  string `gorm:"type:blob"`
	MatchesDuo      string `gorm:"type:blob"`
	MatchesDuoFPP   string `gorm:"type:blob"`
	MatchesSquad    string `gorm:"type:blob"`
	MatchesSquadFPP string `gorm:"type:blob"`
}

func (m *ModelLifeTimeStats) ParseFromPUBG(data schema.LifeTimeStats) {
	m.AccountId = data.Player.ID
	m.GameModeStats = utils.QuickMarshal(data.GameModeStats)
	m.MatchesSolo = utils.QuickMarshal(data.MatchesSolo)
	m.MatchesSoloFPP = utils.QuickMarshal(data.MatchesSoloFPP)
	m.MatchesDuo = utils.QuickMarshal(data.MatchesDuo)
	m.MatchesDuoFPP = utils.QuickMarshal(data.MatchesDuoFPP)
	m.MatchesSquad = utils.QuickMarshal(data.MatchesSquad)
	m.MatchesSquadFPP = utils.QuickMarshal(data.MatchesSquadFPP)
}

func (m *ModelLifeTimeStats) GetInfoByAccountId(shard api.ShardType, accountId string) error {
	// 如果上次数据库更新时间在5分钟内 则直接返回数据库中的数据
	err := db.GetDb().Where("account_id = ? and TIME_TO_SEC(TIMEDIFF(now(), updated_at)) < 60 * 5", accountId).First(m).Error
	if err != nil {
		// 	如果数据库中没有数据的话 发送请求
		remoteStats, err := pubg.Api.RequestLifeTimeStats(shard, accountId)
		if err != nil {
			return nil
		}

		m.ParseFromPUBG(*remoteStats)

		db.GetDb().Where("account_id = ?", m.AccountId).Delete(&ModelLifeTimeStats{})
		db.GetDb().Create(m)
	}

	return nil
}

func (m *ModelLifeTimeStats) GetInfoByUserName(shard api.ShardType, username string) error {
	player := new(ModelPlayer)
	var err error
	err = player.GetInfoByUserName(shard, username)
	if err != nil {
		return err
	}

	err = m.GetInfoByAccountId(shard, player.AccountId)
	if err != nil {
		return err
	}

	return nil
}
