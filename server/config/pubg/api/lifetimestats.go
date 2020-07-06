package api

import (
	"fmt"
	"github.com/moonrailgun/PUBGo/server/config/pubg/schema"
)

func (a *API) RequestLifeTimeStats(accountId string) (*schema.LifeTimeStats, error) {
	endpointUrl := fmt.Sprintf("%s/players/%s/seasons/lifetime", a.Url, accountId)

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
