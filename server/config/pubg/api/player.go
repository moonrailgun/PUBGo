package api

import (
	"fmt"
	"github.com/moonrailgun/PUBGo/server/config/pubg/schema"
	"net/url"
)

func (a *API) RequestSinglePlayerByName(shard ShardType, playerName string) (*schema.Player, error) {
	parameters := url.Values{
		"filter[playerNames]": {playerName},
	}

	endpointUrl := fmt.Sprintf("%s/shards/%s/players?%s", a.Url, shard, parameters.Encode())

	buffer, err := httpRequest(endpointUrl, a.Key); if err != nil {
		return nil, err
	}

	players, err:= schema.ParsePlayers(buffer.String()); if err != nil {
		return nil, err
	}

	return players[0], nil
}