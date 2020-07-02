package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moonrailgun/PUBGo/server/routes/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	player := new(controllers.PlayerController)
	r.GET("/player/shard/:shard/name/:username", player.GetPlayerInfo)

	match := new(controllers.MatchController)
	r.GET("/match/shard/:shard/match/:matchId", match.GetMatchInfo)

	return r
}
