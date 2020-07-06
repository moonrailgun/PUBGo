package pubg

import (
	"fmt"
	"testing"

	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
	"github.com/moonrailgun/PUBGo/server/config/pubg/schema"
	"github.com/stretchr/testify/assert"
)

func TestRequestSinglePlayerByName(t *testing.T) {
	playerName := "WackyJacky101"
	player, err := Api.RequestSinglePlayerByName(api.STEAM, playerName); if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, player.Name, playerName)

	json, err := schema.StringifyPlayer(player); if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json)
}
