package models

import (
	"github.com/jinzhu/gorm"
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
