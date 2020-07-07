package schema

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testLifetimestatsStr = `
{
  "data": {
    "type": "playerSeason",
    "attributes": {
      "gameModeStats": {
        "duo": {
          "assists": 0,
          "boosts": 0,
          "dBNOs": 0,
          "dailyKills": 0,
          "dailyWins": 0,
          "damageDealt": 0,
          "days": 0,
          "headshotKills": 0,
          "heals": 0,
          "killPoints": 0,
          "kills": 0,
          "longestKill": 0,
          "longestTimeSurvived": 0,
          "losses": 0,
          "maxKillStreaks": 0,
          "mostSurvivalTime": 0,
          "rankPoints": 0,
          "rankPointsTitle": "0",
          "revives": 0,
          "rideDistance": 0,
          "roadKills": 0,
          "roundMostKills": 0,
          "roundsPlayed": 0,
          "suicides": 0,
          "swimDistance": 0,
          "teamKills": 0,
          "timeSurvived": 0,
          "top10s": 0,
          "vehicleDestroys": 0,
          "walkDistance": 0,
          "weaponsAcquired": 0,
          "weeklyKills": 0,
          "weeklyWins": 0,
          "winPoints": 0,
          "wins": 0
        },
        "duo-fpp": {
          "assists": 0,
          "boosts": 0,
          "dBNOs": 0,
          "dailyKills": 0,
          "dailyWins": 0,
          "damageDealt": 0,
          "days": 0,
          "headshotKills": 0,
          "heals": 0,
          "killPoints": 0,
          "kills": 0,
          "longestKill": 0,
          "longestTimeSurvived": 0,
          "losses": 0,
          "maxKillStreaks": 0,
          "mostSurvivalTime": 0,
          "rankPoints": 0,
          "rankPointsTitle": "0",
          "revives": 0,
          "rideDistance": 0,
          "roadKills": 0,
          "roundMostKills": 0,
          "roundsPlayed": 0,
          "suicides": 0,
          "swimDistance": 0,
          "teamKills": 0,
          "timeSurvived": 0,
          "top10s": 0,
          "vehicleDestroys": 0,
          "walkDistance": 0,
          "weaponsAcquired": 0,
          "weeklyKills": 0,
          "weeklyWins": 0,
          "winPoints": 0,
          "wins": 0
        },
        "solo": {
          "assists": 5,
          "boosts": 80,
          "dBNOs": 0,
          "dailyKills": 0,
          "dailyWins": 0,
          "damageDealt": 9659.662,
          "days": 43,
          "headshotKills": 19,
          "heals": 111,
          "killPoints": 0,
          "kills": 78,
          "longestKill": 384.85153,
          "longestTimeSurvived": 1871.184,
          "losses": 104,
          "maxKillStreaks": 2,
          "mostSurvivalTime": 1871.184,
          "rankPoints": 0,
          "rankPointsTitle": "0",
          "revives": 0,
          "rideDistance": 309226.3,
          "roadKills": 1,
          "roundMostKills": 9,
          "roundsPlayed": 105,
          "suicides": 7,
          "swimDistance": 1419.9777,
          "teamKills": 7,
          "timeSurvived": 64905.168,
          "top10s": 4,
          "vehicleDestroys": 6,
          "walkDistance": 73691.21,
          "weaponsAcquired": 258,
          "weeklyKills": 0,
          "weeklyWins": 0,
          "winPoints": 0,
          "wins": 1
        },
        "solo-fpp": {
          "assists": 0,
          "boosts": 0,
          "dBNOs": 0,
          "dailyKills": 0,
          "dailyWins": 0,
          "damageDealt": 0,
          "days": 0,
          "headshotKills": 0,
          "heals": 0,
          "killPoints": 0,
          "kills": 0,
          "longestKill": 0,
          "longestTimeSurvived": 0,
          "losses": 0,
          "maxKillStreaks": 0,
          "mostSurvivalTime": 0,
          "rankPoints": 0,
          "rankPointsTitle": "0",
          "revives": 0,
          "rideDistance": 0,
          "roadKills": 0,
          "roundMostKills": 0,
          "roundsPlayed": 0,
          "suicides": 0,
          "swimDistance": 0,
          "teamKills": 0,
          "timeSurvived": 0,
          "top10s": 0,
          "vehicleDestroys": 0,
          "walkDistance": 0,
          "weaponsAcquired": 0,
          "weeklyKills": 0,
          "weeklyWins": 0,
          "winPoints": 0,
          "wins": 0
        },
        "squad": {
          "assists": 898,
          "boosts": 3207,
          "dBNOs": 2068,
          "dailyKills": 4,
          "dailyWins": 0,
          "damageDealt": 321533.25,
          "days": 313,
          "headshotKills": 446,
          "heals": 5585,
          "killPoints": 0,
          "kills": 2182,
          "longestKill": 441.83667,
          "longestTimeSurvived": 2017.243,
          "losses": 2437,
          "maxKillStreaks": 4,
          "mostSurvivalTime": 2017.243,
          "rankPoints": 0,
          "rankPointsTitle": "0",
          "revives": 633,
          "rideDistance": 1386682.8,
          "roadKills": 6,
          "roundMostKills": 10,
          "roundsPlayed": 2525,
          "suicides": 43,
          "swimDistance": 7490.3784,
          "teamKills": 50,
          "timeSurvived": 1662879.4,
          "top10s": 783,
          "vehicleDestroys": 13,
          "walkDistance": 2759651.2,
          "weaponsAcquired": 9822,
          "weeklyKills": 12,
          "weeklyWins": 1,
          "winPoints": 0,
          "wins": 93
        },
        "squad-fpp": {
          "assists": 14,
          "boosts": 24,
          "dBNOs": 10,
          "dailyKills": 0,
          "dailyWins": 0,
          "damageDealt": 2313.5137,
          "days": 8,
          "headshotKills": 1,
          "heals": 45,
          "killPoints": 0,
          "kills": 15,
          "longestKill": 54.04946,
          "longestTimeSurvived": 1379.455,
          "losses": 35,
          "maxKillStreaks": 1,
          "mostSurvivalTime": 1379.455,
          "rankPoints": 0,
          "rankPointsTitle": "0",
          "revives": 9,
          "rideDistance": 7127.345,
          "roadKills": 0,
          "roundMostKills": 2,
          "roundsPlayed": 38,
          "suicides": 0,
          "swimDistance": 0,
          "teamKills": 0,
          "timeSurvived": 17522.625,
          "top10s": 12,
          "vehicleDestroys": 0,
          "walkDistance": 33419.48,
          "weaponsAcquired": 111,
          "weeklyKills": 0,
          "weeklyWins": 0,
          "winPoints": 0,
          "wins": 4
        }
      },
      "bestRankPoint": 0
    },
    "relationships": {
      "matchesDuo": {
        "data": []
      },
      "matchesDuoFPP": {
        "data": []
      },
      "matchesSquad": {
        "data": [
          {
            "type": "match",
            "id": "af48c404-08e1-4822-a2e9-fa1a4fbb3980"
          },
          {
            "type": "match",
            "id": "a72a9164-3291-4397-b752-dca04f58b100"
          },
          {
            "type": "match",
            "id": "186f4ee7-235c-4778-b581-8114821b8e05"
          },
          {
            "type": "match",
            "id": "4e15e5c1-37af-48e6-9f09-ab62c2d756b1"
          },
          {
            "type": "match",
            "id": "17a19fdf-b95e-4fd0-b041-3720afd789e5"
          },
          {
            "type": "match",
            "id": "44deace8-a1bf-4ee9-88bc-10b60b340313"
          },
          {
            "type": "match",
            "id": "3271267a-1bba-4c5f-a33c-68de7c2646d5"
          },
          {
            "type": "match",
            "id": "33d584e4-2cc3-41bf-8ef1-e7b88c0aa0c4"
          },
          {
            "type": "match",
            "id": "835d0852-7000-47fb-babe-6efcab09f161"
          },
          {
            "type": "match",
            "id": "7c3736ba-599f-4803-a1de-1a55c8634eee"
          },
          {
            "type": "match",
            "id": "56b80767-d090-46bc-9667-c21ac5ca110a"
          },
          {
            "type": "match",
            "id": "b8482ef3-3fff-4a64-85fa-a81254f797c9"
          },
          {
            "type": "match",
            "id": "1225d3b3-d060-4975-954d-72b43f5ae754"
          },
          {
            "type": "match",
            "id": "2a244d5a-9601-4060-bbbc-b43cd571ab24"
          },
          {
            "type": "match",
            "id": "075b2c60-e305-4cf7-b95c-c084248fbcac"
          },
          {
            "type": "match",
            "id": "2bbb220c-69e6-4bd6-a4bc-cd3052872e11"
          },
          {
            "type": "match",
            "id": "6c91ba99-21c8-4e60-8b21-64ec152a28cc"
          },
          {
            "type": "match",
            "id": "133c67f6-2b92-410f-bf01-5b5a6922cdb0"
          },
          {
            "type": "match",
            "id": "c08c4a8b-0173-46cc-9cd1-75206d64a184"
          },
          {
            "type": "match",
            "id": "b6ca84db-913a-40e5-8d05-788871b54882"
          },
          {
            "type": "match",
            "id": "73740aa0-ecf8-4cf8-ab1a-825ac776693a"
          },
          {
            "type": "match",
            "id": "eadd7ef6-e163-4995-9308-ac555ee33b86"
          },
          {
            "type": "match",
            "id": "3248d2dd-d199-4107-b283-dbf530430153"
          },
          {
            "type": "match",
            "id": "34bb767e-5e83-4395-836c-df56b5b803ab"
          }
        ]
      },
      "matchesSquadFPP": {
        "data": []
      },
      "season": {
        "data": {
          "type": "season",
          "id": "lifetime"
        }
      },
      "player": {
        "data": {
          "type": "player",
          "id": "account.5c3d60147eb549d18a0984a67af3f58b"
        }
      },
      "matchesSolo": {
        "data": []
      },
      "matchesSoloFPP": {
        "data": []
      }
    }
  },
  "links": {
    "self": "https://api.pubg.com/shards/steam/players/account.5c3d60147eb549d18a0984a67af3f58b/seasons/lifetime"
  },
  "meta": {}
}
`

func TestParseParseLifeTimeStats(t *testing.T) {
	stats, err := ParseLifeTimeStats(testLifetimestatsStr)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 14, stats.GameModeStats.SquadFpp.Assists)
	assert.Equal(t, 24, stats.GameModeStats.SquadFpp.Boosts)
	assert.Greater(t, len(stats.MatchesSquad), 0)
	assert.Equal(t, "af48c404-08e1-4822-a2e9-fa1a4fbb3980", stats.MatchesSquad[0].ID)
}
