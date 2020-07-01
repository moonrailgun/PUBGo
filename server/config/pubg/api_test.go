package pubg

import (
	"fmt"
	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
	"github.com/moonrailgun/PUBGo/server/config/pubg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestSinglePlayerByName(t *testing.T) {
	playerName := "WackyJacky101"
	player, err := Api.RequestSinglePlayerByName(api.STEAM, playerName); if err != nil {
		panic(err)
	}

	assert.Equal(t, player.Name, playerName)

	//fmt.Printf("%+v", player)

	json, err := model.StringifyPlayer(player); if err != nil {
		panic(err)
	}
	fmt.Println(json)
}