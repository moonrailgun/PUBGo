package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/moonrailgun/PUBGo/server/config/pubg/schema"
	"time"
)

type ModelMatch struct {
	gorm.Model
	MatchId        string `gorm:"not null;unique_index"`
	Duration       int
	MatchCreatedAt time.Time
	Roster         string `gorm:"type:blob"`
}

func (m *ModelMatch) ParseFromPUBG(data schema.Match) {
	m.MatchId = data.ID
	m.Duration = data.Duration
	m.MatchCreatedAt = data.CreatedAt

	roster, _ := json.Marshal(data.Rosters)
	m.Roster = string(roster[:])
}
