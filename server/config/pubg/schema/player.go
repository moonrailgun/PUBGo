package schema

import (
	"bytes"
	"errors"
	"reflect"
	"time"

	"github.com/google/jsonapi"
)

type Player struct {
	ID           string          `jsonapi:"primary,player"`
	Name         string          `jsonapi:"attr,name"`
	ShardID      string          `jsonapi:"attr,shardId"`
	CreatedAt    time.Time       `jsonapi:"attr,createdAt,iso8601"`
	UpdatedAt    time.Time       `jsonapi:"attr,updatedAt,iso8601"`
	PatchVersion string          `jsonapi:"attr,patchVersion"`
	TitleID      string          `jsonapi:"attr,titleId"`
	Matches      []*MatchSummary `jsonapi:"relation,matches"`
}

type PlayerSummary struct {
	ID string `jsonapi:"primary,player"`
}

func ParsePlayers(jsonStr string) ([]*Player, error) {
	in := bytes.NewReader([]byte(jsonStr))
	result, err := jsonapi.UnmarshalManyPayload(in, reflect.TypeOf(new(Player)))
	if err != nil {
		return nil, err
	}

	players := make([]*Player, len(result))
	for idx, elt := range result {
		player, ok := elt.(*Player)
		if !ok {
			return nil, errors.New("Failed to convert players. ")
		}
		players[idx] = player
	}
	return players, nil
}

func StringifyPlayer(player *Player) (string, error) {
	var s bytes.Buffer
	err := jsonapi.MarshalPayload(&s, player)
	if err != nil {
		return "", err
	}

	return s.String(), nil
}
