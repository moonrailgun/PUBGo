package models

import (
	"github.com/moonrailgun/PUBGo/server/config/db"
	"github.com/moonrailgun/PUBGo/server/config/pubg"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
	"github.com/moonrailgun/PUBGo/server/config/pubg/schema"
	"github.com/moonrailgun/PUBGo/server/utils"
	"time"
)

type Match struct {
	BaseModel
	MatchId        string `gorm:"not null;unique_index"`
	Duration       int
	MatchType      string // arcade, custom, event, official, training
	MapName        string // Baltic_Main, Desert_Main, DihorOtok_Main, Erangel_Main, Range_Main, Savage_Main, Summerland_Main
	IsCustomMatch  bool
	SeasonState    string // closed, prepare, progress
	MatchCreatedAt time.Time
	Roster         string `gorm:"type:blob"`
	Assets         string `gorm:"type:blob"`
}

func (m *Match) ParseFromPUBG(data schema.Match) {
	m.MatchId = data.ID
	m.Duration = data.Duration
	m.MatchType = data.MatchType
	m.MapName = data.MapName
	m.IsCustomMatch = data.IsCustomMatch
	m.SeasonState = data.SeasonState
	m.MatchCreatedAt = data.CreatedAt

	m.Roster = utils.QuickMarshal(data.Rosters)
	m.Assets = utils.QuickMarshal(data.Assets)
}

// 获取比赛详情
// 并进行数据库缓存
func (m *Match) GetMatchInfo(shard api.ShardType, matchId string) error {
	err := db.GetDb().Where("match_id = ?", matchId).First(m).Error
	if err != nil {
		// 如果没有取到数据的话
		remoteMatch, err := pubg.Api.RequestMatchInfo(api.ShardType(shard), matchId)
		if err != nil {
			return err
		}

		m.ParseFromPUBG(*remoteMatch)
		db.GetDb().Create(m)
	}

	return nil
}
