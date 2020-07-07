package pubg

import (
	"fmt"
	"testing"

	"github.com/moonrailgun/PUBGo/server/config/pubg/api"
	"github.com/moonrailgun/PUBGo/server/config/pubg/schema"
	"github.com/stretchr/testify/assert"
)

func TestRequestSinglePlayerByName(t *testing.T) {
	testPlayerName := "WackyJacky101"
	player, err := Api.RequestSinglePlayerByName(api.STEAM, testPlayerName); if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, player.Name, testPlayerName)

	json, err := schema.StringifyPlayer(player); if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json)
}
