package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moonrailgun/PUBGo/server/routes/controllers"
	stats "github.com/semihalev/gin-stats"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(stats.RequestStats())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})

	player := new(controllers.PlayerController)
	r.GET("/player/shard/:shard/name/:username", player.GetPlayerInfo)

	match := new(controllers.MatchController)
	r.GET("/match/shard/:shard/match/:matchId", match.GetMatchInfo)

	return r
}
