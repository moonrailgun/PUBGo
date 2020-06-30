package pubg

import (
	"fmt"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
	"testing"
)

func TestRequestSinglePlayerByName(t *testing.T) {
	player, err := Api.RequestSinglePlayerByName(api.STEAM, "WackyJacky101"); if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", player)
}