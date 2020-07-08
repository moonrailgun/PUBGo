package models

import (
	"github.com/moonrailgun/PUBGo/server/config/db"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayerGetInfoByUserName(t *testing.T) {
	testPlayerName := "WackyJacky101"

	player := new(ModelPlayer)
	err := player.GetInfoByUserName(api.STEAM, testPlayerName)
	if err != nil {
		panic(err)
	}

	// 检查内存
	assert.Equal(t, testPlayerName, player.Name)
	assert.Equal(t, "account.c0e530e9b7244b358def282782f893af", player.AccountId)
	assert.Equal(t, "steam", player.ShardId)

	// 检查是否已经写入数据库
	dbPlayer := new(ModelPlayer)
	db.GetDb().Where("id = ?", player.ID).First(dbPlayer)
	assert.Equal(t, testPlayerName, dbPlayer.Name)
}
