package api

import (
	"fmt"
	"github.com/moonrailgun/PUBGo/server/config/pubg/schema"
)

func (a *API) RequestLifeTimeStats(shard ShardType, accountId string) (*schema.LifeTimeStats, error) {
	endpointUrl := fmt.Sprintf("%s/shards/%s/players/%s/seasons/lifetime", a.Url, shard, accountId)

	buffer, err := httpRequest(endpointUrl, a.Key)
	if err != nil {
		return nil, err
	}

	stats, err := schema.ParseLifeTimeStats(buffer.String())
	if err != nil {
		return nil, err
	}

	return stats, nil
}
