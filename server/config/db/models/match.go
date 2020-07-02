package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ModelMatch struct {
	gorm.Model
	MatchId        string `gorm:"not null;unique_index"`
	Duration       int
	MatchCreatedAt time.Time
}
