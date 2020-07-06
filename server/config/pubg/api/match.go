package api

import (
	"fmt"
	"github.com/moonrailgun/PUBGo/server/config/pubg/schema"
)

func (a *API) RequestMatchInfo(shard ShardType, matchId string) (*schema.Match, error) {
	endpointUrl := fmt.Sprintf("%s/shards/%s/matches/%s", a.Url, shard, matchId)

	buffer, err := httpRequest(endpointUrl, a.Key)
	if err != nil {
		return nil, err
	}

	match, err := schema.ParseMatch(buffer.String())
	if err != nil {
		return nil, err
	}

	return match, nil
}
