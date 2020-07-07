package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/moonrailgun/PUBGo/server/config/db/models"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
)

type StatsController struct{}

func (ctl StatsController) GetLifeTimeStats(c *gin.Context) {
	shard := c.Param("shard")
	username := c.Param("username")

	stats := new(models.ModelLifeTimeStats)
	err := stats.GetInfoByUserName(api.ShardType(shard), username)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}

	c.JSON(200, stats)
}
