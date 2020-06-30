package pubg

import (
	"github.com/moonrailgun/PUBGo/server/config"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
)
var Api *api.API

func init() {
	Api = api.NewAPI(config.Config.PUBG.API)
}
