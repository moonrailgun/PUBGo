package models

import (
	"time"
)

type PubgRequestLog struct {
	ID        uint `json:"id" gorm:"primary_key"`
	Url       string
	Body      string
	Resp      string `gorm:"type:blob"`
	CreatedAt time.Time
}
