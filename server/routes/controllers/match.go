package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/moonrailgun/PUBGo/server/config/db"
	"github.com/moonrailgun/PUBGo/server/config/db/models"
	"github.com/moonrailgun/PUBGo/server/config/pubg"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
)

type MatchController struct{}

func (ctl MatchController) GetMatchInfo(c *gin.Context) {
	shard := c.Param("shard")
	matchId := c.Param("matchId")

	// 优先从数据库中查找
	var match models.ModelMatch
	err := db.GetDb().Where("match_id = ?", matchId).First(&match).Error
	if err != nil {
		// 如果没有取到数据的话
		remoteMatch, err := pubg.Api.RequestMatchInfo(api.ShardType(shard), matchId)
		if err != nil {
			c.JSON(500, gin.H{"msg": err})
			return
		}

		match.ParseFromPUBG(*remoteMatch)
		db.GetDb().Create(&match)
	}

	c.JSON(200, match)
}
