package schema

import (
	"bytes"
	"errors"
	"github.com/google/jsonapi"
	"reflect"
)

type Season struct {
	SeasonId        string `jsonapi:"primary,season"`
	IsCurrentSeason bool   `jsonapi:"attr,isCurrentSeason"`
	IsOffseason     bool   `jsonapi:"attr,isOffseason"`
}

func ParseSeasons(jsonStr string) ([]*Season, error) {
	in := bytes.NewReader([]byte(jsonStr))
	result, err := jsonapi.UnmarshalManyPayload(in, reflect.TypeOf(new(Season)))
	if err != nil {
		return nil, err
	}

	seasons := make([]*Season, len(result))
	for idx, elt := range result {
		season, ok := elt.(*Season)
		if !ok {
			return nil, errors.New("Failed to convert seasons. ")
		}
		seasons[idx] = season
	}

	return seasons, nil
}
