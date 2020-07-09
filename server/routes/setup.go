package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moonrailgun/PUBGo/server/routes/controllers"
	ginStats "github.com/semihalev/gin-stats"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(ginStats.RequestStats())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, ginStats.Report())
	})

	player := new(controllers.PlayerController)
	r.GET("/player/shard/:shard/name/:username", player.GetPlayerInfo)

	match := new(controllers.MatchController)
	r.GET("/match/shard/:shard/match/:matchId", match.GetMatchInfo)

	stats := new(controllers.StatsController)
	r.GET("/stats/shard/:shard/name/:username", stats.GetLifeTimeStats)
	r.GET("/stats/shard/:shard/name/:username/renew", stats.GetLifeTimeStatsRenew)

	return r
}
