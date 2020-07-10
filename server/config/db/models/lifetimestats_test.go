package models

import (
	"github.com/moonrailgun/PUBGo/server/config/db"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLifeTimeStatsGetInfoByUserName(t *testing.T) {
	testPlayerName := "WackyJacky101"

	stats := new(LifeTimeStats)
	err := stats.GetInfoByUserName(api.STEAM, testPlayerName, false)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "account.c0e530e9b7244b358def282782f893af", stats.AccountId)
	assert.False(t, db.GetDb().NewRecord(stats))
}
