package api

import (
	"fmt"
	"github.com/moonrailgun/PUBGo/server/config/pubg/model"
)

func (a *API) RequestMatchInfo(shard ShardType, matchId string) (*model.Match, error) {
	endpointUrl := fmt.Sprintf("%s/shards/%s/matches/%s", a.Url, shard, matchId)

	buffer, err := httpRequest(endpointUrl, a.Key)
	if err != nil {
		return nil, err
	}

	match, err := model.ParseMatch(buffer.String())
	if err != nil {
		return nil, err
	}

	return match, nil
}
