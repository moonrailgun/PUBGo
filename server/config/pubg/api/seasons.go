package api

import (
	"fmt"
	"github.com/moonrailgun/PUBGo/server/config/pubg/schema"
)

func (a *API) RequestSeasons(shard ShardType) ([]*schema.Season, error) {
	endpointUrl := fmt.Sprintf("%s/shards/%s/seasons", a.Url, shard)

	buffer, err := httpRequest(endpointUrl, a.Key)
	if err != nil {
		return nil, err
	}

	seasons, err := schema.ParseSeasons(buffer.String())
	if err != nil {
		return nil, err
	}

	return seasons, nil
}
