package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/moonrailgun/PUBGo/server/config/db"
	"github.com/moonrailgun/PUBGo/server/config/db/models"
	"github.com/moonrailgun/PUBGo/server/config/pubg"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
)

type PlayerController struct{}

// 获取用户信息
func (ctl PlayerController) GetPlayerInfo(c *gin.Context) {
	shard := c.Param("shard")
	username := c.Param("username")

	remotePlayer, err := pubg.Api.RequestSinglePlayerByName(api.ShardType(shard), username)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}

	// 将账号信息存储到数据库
	var player models.ModelPlayer
	player.ParseFromPUBG(*remotePlayer)
	db.GetDb().Where("account_id = ?", remotePlayer.ID).FirstOrCreate(&player)

	c.JSON(200, remotePlayer)
}
