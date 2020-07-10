package schema

import (
	"bytes"
	"github.com/google/jsonapi"
	"time"
)

// Match structure represent data related to a PUBG match
type MatchSummary struct {
	ID string `jsonapi:"primary,match" json:"id"`
	// GameID string `jsonapi:"attr,id"`
}

type Match struct {
	ID            string    `jsonapi:"primary,match" json:"id"`
	CreatedAt     time.Time `jsonapi:"attr,createdAt,iso8601"`
	Duration      int       `jsonapi:"attr,duration"`
	GameMode      string    `jsonapi:"attr,gameMode"`
	MatchType     string    `jsonapi:"attr,matchType"` // arcade, custom, event, official, training
	MapName       string    `jsonapi:"attr,mapName"`   // Baltic_Main, Desert_Main, DihorOtok_Main, Erangel_Main, Range_Main, Savage_Main, Summerland_Main
	IsCustomMatch bool      `jsonapi:"attr,isCustomMatch"`
	PatchVersion  string    `jsonapi:"attr,patchVersion"`
	SeasonState   string    `jsonapi:"attr,seasonState"` // closed, prepare, progress
	ShardId       string    `jsonapi:"attr,shardId"`
	TitleId       string    `jsonapi:"attr,titleId"`
	Rosters       []*Roster `jsonapi:"relation,rosters"`
	Assets        []*Asset  `jsonapi:"relation,assets"`
}

type Asset struct {
	ID          string    `jsonapi:"primary,asset" json:"id"`
	Url         string    `jsonapi:"attr,URL"`
	Name        string    `jsonapi:"attr,name"`
	Description string    `jsonapi:"attr,description"`
	CreatedAt   time.Time `jsonapi:"attr,createdAt,iso8601"`
}

type Roster struct {
	ID      string `jsonapi:"primary,roster" json:"id"`
	ShardId string `jsonapi:"attr,shardId"`
	Stats   struct {
		Rank   int `jsonapi:"attr,rank"`
		TeamID int `jsonapi:"attr,teamId"`
	} `jsonapi:"attr,stats"`
	Won          string         `jsonapi:"attr,won"`
	Participants []*Participant `jsonapi:"relation,participants"`
}

type Participant struct {
	ID      string `jsonapi:"primary,participant" json:"id"`
	Actor   string `jsonapi:"attr,actor"`
	ShardId string `jsonapi:"attr,shardId"`
	Stats   struct {
		DBNOs           int     `jsonapi:"attr,DBNOs"`
		Assists         int     `jsonapi:"attr,assists"`
		Boosts          int     `jsonapi:"attr,boosts"`
		DamageDealt     float64 `jsonapi:"attr,damageDealt"`
		DeathType       string  `jsonapi:"attr,deathType"`
		HeadshotKills   int     `jsonapi:"attr,headshotKills"`
		Heals           int     `jsonapi:"attr,heals"`
		KillPlace       int     `jsonapi:"attr,killPlace"`
		KillPoints      int     `jsonapi:"attr,killPoints"`
		KillPointsDelta float64 `jsonapi:"attr,killPointsDelta"`
		KillStreaks     int     `jsonapi:"attr,killStreaks"`
		Kills           int     `jsonapi:"attr,kills"`
		LastKillPoints  int     `jsonapi:"attr,lastKillPoints"`
		LastWinPoints   int     `jsonapi:"attr,lastWinPoints"`
		LongestKill     int     `jsonapi:"attr,longestKill"`
		MostDamage      int     `jsonapi:"attr,mostDamage"`
		Name            string  `jsonapi:"attr,name"`
		PlayerID        string  `jsonapi:"attr,playerId"`
		Revives         int     `jsonapi:"attr,revives"`
		RideDistance    float64 `jsonapi:"attr,rideDistance"`
		RoadKills       int     `jsonapi:"attr,roadKills"`
		TeamKills       int     `jsonapi:"attr,teamKills"`
		TimeSurvived    float64 `jsonapi:"attr,timeSurvived"`
		VehicleDestroys int     `jsonapi:"attr,vehicleDestroys"`
		WalkDistance    float64 `jsonapi:"attr,walkDistance"`
		WeaponsAcquired int     `jsonapi:"attr,weaponsAcquired"`
		WinPlace        int     `jsonapi:"attr,winPlace"`
		WinPoints       int     `jsonapi:"attr,winPoints"`
		WinPointsDelta  float64 `jsonapi:"attr,winPointsDelta"`
	} `jsonapi:"attr,stats"`
}

func ParseMatch(jsonStr string) (*Match, error) {
	in := bytes.NewReader([]byte(jsonStr))
	match := new(Match)
	err := jsonapi.UnmarshalPayload(in, match)
	if err != nil {
		return nil, err
	}

	return match, nil
}
