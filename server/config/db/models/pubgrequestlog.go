package models

import (
	"time"
)

type ModelPubgRequestLog struct {
	ID        uint `gorm:"primary_key"`
	Url       string
	Body      string
	Resp      string `gorm:"type:blob"`
	CreatedAt time.Time
}
