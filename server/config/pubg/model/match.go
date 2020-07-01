package model

import "time"

type Match struct {
	ID           string    `jsonapi:"primary,match"`
	CreatedAt    time.Time `jsonapi:"attr,createdAt,iso8601"`
	Duration     int       `jsonapi:"attr,duration"`
	GameMode     string    `jsonapi:"attr,gameMode"`
	PatchVersion string    `jsonapi:"attr,patchVersion"`
	ShardID      string    `jsonapi:"attr,shardId"`
	TitleID      string    `jsonapi:"attr,titleId"`
	Rosters      []*Roster `jsonapi:"relation,rosters"`
}

type Roster struct {
	ID      string `jsonapi:"primary,roster"`
	ShardID string `jsonapi:"attr,shardId"`
	Stats   struct {
		Rank   int `json:"rank"`
		TeamID int `json:"teamId"`
	} `jsonapi:"attr,stats"`
	Won          string         `jsonapi:"attr,won"`
	Participants []*Participant `jsonapi:"relation,participants"`
}

type Participant struct {
	ID      string `jsonapi:"primary,participant"`
	Actor   string `jsonapi:"attr,actor"`
	ShardID string `jsonapi:"attr,shardId"`
	Stats   struct {
		DBNOs           int     `json:"DBNOs"`
		Assists         int     `json:"assists"`
		Boosts          int     `json:"boosts"`
		DamageDealt     float64 `json:"damageDealt"`
		DeathType       string  `json:"deathType"`
		HeadshotKills   int     `json:"headshotKills"`
		Heals           int     `json:"heals"`
		KillPlace       int     `json:"killPlace"`
		KillPoints      int     `json:"killPoints"`
		KillPointsDelta float64 `json:"killPointsDelta"`
		KillStreaks     int     `json:"killStreaks"`
		Kills           int     `json:"kills"`
		LastKillPoints  int     `json:"lastKillPoints"`
		LastWinPoints   int     `json:"lastWinPoints"`
		LongestKill     int     `json:"longestKill"`
		MostDamage      int     `json:"mostDamage"`
		Name            string  `json:"name"`
		PlayerID        string  `json:"playerId"`
		Revives         int     `json:"revives"`
		RideDistance    float64 `json:"rideDistance"`
		RoadKills       int     `json:"roadKills"`
		TeamKills       int     `json:"teamKills"`
		TimeSurvived    float64 `json:"timeSurvived"`
		VehicleDestroys int     `json:"vehicleDestroys"`
		WalkDistance    float64 `json:"walkDistance"`
		WeaponsAcquired int     `json:"weaponsAcquired"`
		WinPlace        int     `json:"winPlace"`
		WinPoints       int     `json:"winPoints"`
		WinPointsDelta  float64 `json:"winPointsDelta"`
	} `jsonapi:"attr,stats"`
}
