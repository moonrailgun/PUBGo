package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/moonrailgun/PUBGo/server/config/pubg"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
)

type MatchController struct{}

func (ctl MatchController) GetMatchInfo(c *gin.Context) {
	shard := c.Param("shard")
	matchId := c.Param("matchId")
	match, err := pubg.Api.RequestMatchInfo(api.ShardType(shard), matchId)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
	}

	c.JSON(200, match)
}
