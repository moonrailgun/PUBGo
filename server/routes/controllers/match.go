package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/moonrailgun/PUBGo/server/config/db/models"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
)

type MatchController struct{}

func (ctl MatchController) GetMatchInfo(c *gin.Context) {
	shard := c.Param("shard")
	matchId := c.Param("matchId")

	// 优先从数据库中查找
	var match models.Match
	err := match.GetMatchInfo(api.ShardType(shard), matchId)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}

	c.JSON(200, match)
}
