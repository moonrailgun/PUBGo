package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/moonrailgun/PUBGo/server/config/pubg"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
)

type PlayerController struct{}

// 获取用户信息
func (ctl PlayerController) GetPlayerInfo(c *gin.Context) {
	shard := c.Param("shard")
	username := c.Param("username")
	player, err := pubg.Api.RequestSinglePlayerByName(api.ShardType(shard), username)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
	}

	c.JSON(200, player)
}
