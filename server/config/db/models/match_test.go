package models

import (
	"github.com/moonrailgun/PUBGo/server/config/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMatchInfo(t *testing.T) {
	testMatchId := "186f4ee7-235c-4778-b581-8114821b8e05"

	match := new(ModelMatch)
	err := match.GetMatchInfo("steam", testMatchId)
	if err != nil {
		panic(err)
	}

	// 检查内存
	assert.False(t, db.GetDb().NewRecord(match)) // 判断已经被写入
	assert.Equal(t, testMatchId, match.MatchId)
	assert.Greater(t, match.ID, uint(0))

	// 检查是否已经写入数据库
	dbMatch := new(ModelMatch)
	db.GetDb().Where("id = ?", match.ID).First(dbMatch)
	assert.Equal(t, testMatchId, dbMatch.MatchId)
}