package schema

import (
	"bytes"
	"github.com/google/jsonapi"
)

type LifeTimeStats struct {
	Player          *PlayerSummary  `jsonapi:"relation,player"`
	GameModeStats   *GameModeStats  `jsonapi:"attr,gameModeStats"`
	BestRankPoint   int             `jsonapi:"attr,bestRankPoint"`
	MatchesSolo     []*MatchSummary `jsonapi:"relation,matchesSolo"`
	MatchesSoloFPP  []*MatchSummary `jsonapi:"relation,matchesSoloFPP"`
	MatchesDuo      []*MatchSummary `jsonapi:"relation,matchesDuo"`
	MatchesDuoFPP   []*MatchSummary `jsonapi:"relation,matchesDuoFPP"`
	MatchesSquad    []*MatchSummary `jsonapi:"relation,matchesSquad"`
	MatchesSquadFPP []*MatchSummary `jsonapi:"relation,matchesSquadFPP"`
}

type GameModeStats struct {
	Duo      *GameModeStatsData `jsonapi:"attr,duo"`
	DuoFPP   *GameModeStatsData `jsonapi:"attr,duo-fpp"`
	Solo     *GameModeStatsData `jsonapi:"attr,solo"`
	SoloFPP  *GameModeStatsData `jsonapi:"attr,solo-fpp"`
	Squad    *GameModeStatsData `jsonapi:"attr,squad"`
	SquadFPP *GameModeStatsData `jsonapi:"attr,squad-fpp"`
}

type GameModeStatsData struct {
	Assists             int     `jsonapi:"attr,assists"`
	Boosts              int     `jsonapi:"attr,boosts"`
	DBNOs               int     `jsonapi:"attr,dBNOs"`
	DailyKills          int     `jsonapi:"attr,dailyKills"`
	DamageDealt         float32 `jsonapi:"attr,damageDealt"`
	Days                int     `jsonapi:"attr,days"`
	DailyWins           int     `jsonapi:"attr,dailyWins"`
	HeadshotKills       int     `jsonapi:"attr,headshotKills"`
	Heals               int     `jsonapi:"attr,heals"`
	KillPoints          int     `jsonapi:"attr,killPoints"`
	Kills               int     `jsonapi:"attr,kills"`
	LongestKill         float32 `jsonapi:"attr,longestKill"`
	LongestTimeSurvived float32 `jsonapi:"attr,longestTimeSurvived"`
	Losses              int     `jsonapi:"attr,losses"`
	MaxKillStreaks      int     `jsonapi:"attr,maxKillStreaks"`
	MostSurvivalTime    float32 `jsonapi:"attr,mostSurvivalTime"`
	RankPoints          int     `jsonapi:"attr,rankPoints"`
	RankPointsTitle     string  `jsonapi:"attr,rankPointsTitle"`
	Revives             int     `jsonapi:"attr,revives"`
	RideDistance        float32 `jsonapi:"attr,rideDistance"`
	RoadKills           int     `jsonapi:"attr,roadKills"`
	RoundMostKills      int     `jsonapi:"attr,roundMostKills"`
	RoundsPlayed        int     `jsonapi:"attr,roundsPlayed"`
	Suicides            int     `jsonapi:"attr,suicides"`
	SwimDistance        float32 `jsonapi:"attr,swimDistance"`
	TeamKills           int     `jsonapi:"attr,teamKills"`
	TimeSurvived        float32 `jsonapi:"attr,timeSurvived"`
	Top10s              int     `jsonapi:"attr,top10s"`
	VehicleDestroys     int     `jsonapi:"attr,vehicleDestroys"`
	WalkDistance        float32 `jsonapi:"attr,walkDistance"`
	WeaponsAcquired     int     `jsonapi:"attr,weaponsAcquired"`
	WeeklyKills         int     `jsonapi:"attr,weeklyKills"`
	WeeklyWins          int     `jsonapi:"attr,weeklyWins"`
	WinPoints           int     `jsonapi:"attr,winPoints"`
	Wins                int     `jsonapi:"attr,wins"`
}

func ParseLifeTimeStats(jsonStr string) (*LifeTimeStats, error) {
	in := bytes.NewReader([]byte(jsonStr))
	lifeTimeStats := new(LifeTimeStats)
	err := jsonapi.UnmarshalPayload(in, lifeTimeStats)
	if err != nil {
		return nil, err
	}

	return lifeTimeStats, nil
}
