package models

import (
	"github.com/moonrailgun/PUBGo/server/config/db"
	"github.com/moonrailgun/PUBGo/server/config/pubg"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
	"github.com/moonrailgun/PUBGo/server/config/pubg/schema"
	"github.com/moonrailgun/PUBGo/server/utils"
	"time"
)

type LifeTimeStats struct {
	ID              uint `json:"id" gorm:"primary_key"`
	UpdatedAt       time.Time
	AccountId       string
	GameModeStats   string `gorm:"type:blob"`
	MatchesSolo     string `gorm:"type:blob"`
	MatchesSoloFPP  string `gorm:"type:blob"`
	MatchesDuo      string `gorm:"type:blob"`
	MatchesDuoFPP   string `gorm:"type:blob"`
	MatchesSquad    string `gorm:"type:blob"`
	MatchesSquadFPP string `gorm:"type:blob"`
}

func (m *LifeTimeStats) ParseFromPUBG(data schema.LifeTimeStats) {
	m.AccountId = data.Player.ID
	m.GameModeStats = utils.QuickMarshal(data.GameModeStats)
	m.MatchesSolo = utils.QuickMarshal(data.MatchesSolo)
	m.MatchesSoloFPP = utils.QuickMarshal(data.MatchesSoloFPP)
	m.MatchesDuo = utils.QuickMarshal(data.MatchesDuo)
	m.MatchesDuoFPP = utils.QuickMarshal(data.MatchesDuoFPP)
	m.MatchesSquad = utils.QuickMarshal(data.MatchesSquad)
	m.MatchesSquadFPP = utils.QuickMarshal(data.MatchesSquadFPP)
}

func (m *LifeTimeStats) GetInfoByAccountId(shard api.ShardType, accountId string, renew bool) error {
	// 如果上次数据库更新时间在5分钟内 则直接返回数据库中的数据

	query := "account_id = ?"
	if renew == true {
		// 如果需要刷新的话则增加条件
		// 当不满足条件的话就直接重新获取
		// 满足条件的话renew不会生效
		query += " and TIME_TO_SEC(TIMEDIFF(now(), updated_at)) < 60 * 5"
	}
	err := db.GetDb().Where(query, accountId).First(m).Error
	if err != nil {
		// 	如果数据库中没有数据的话 发送请求
		remoteStats, err := pubg.Api.RequestLifeTimeStats(shard, accountId)
		if err != nil {
			return nil
		}

		m.ParseFromPUBG(*remoteStats)

		db.GetDb().Where("account_id = ?", m.AccountId).Delete(&LifeTimeStats{})
		db.GetDb().Create(m)
	}

	return nil
}

func (m *LifeTimeStats) GetInfoByUserName(shard api.ShardType, username string, renew bool) error {
	player := new(Player)
	var err error
	err = player.GetInfoByUserName(shard, username)
	if err != nil {
		return err
	}

	err = m.GetInfoByAccountId(shard, player.AccountId, renew)
	if err != nil {
		return err
	}

	return nil
}
