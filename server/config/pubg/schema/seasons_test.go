package schema

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testSeasonsStr = `
{
  "data": [
    {
      "type": "season",
      "id": "division.bro.official.2017-beta",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2017-pre1",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2017-pre2",
      "attributes": {
        "isOffseason": false,
        "isCurrentSeason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2017-pre3",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2017-pre4",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2017-pre5",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2017-pre6",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2017-pre7",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2017-pre8",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2017-pre9",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2018-01",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2018-02",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2018-03",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2018-04",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2018-05",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2018-06",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2018-07",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2018-08",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.2018-09",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.pc-2018-01",
      "attributes": {
        "isOffseason": false,
        "isCurrentSeason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.pc-2018-02",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.pc-2018-03",
      "attributes": {
        "isOffseason": false,
        "isCurrentSeason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.pc-2018-04",
      "attributes": {
        "isOffseason": false,
        "isCurrentSeason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.pc-2018-05",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.pc-2018-06",
      "attributes": {
        "isCurrentSeason": false,
        "isOffseason": false
      }
    },
    {
      "type": "season",
      "id": "division.bro.official.pc-2018-07",
      "attributes": {
        "isCurrentSeason": true,
        "isOffseason": false
      }
    }
  ],
  "links": {
    "self": "https://api.pubg.com/shards/steam/seasons"
  },
  "meta": {}
}
`

func TestParseSeasons(t *testing.T) {
	seasons, err := ParseSeasons(testSeasonsStr)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 26, len(seasons))
	assert.Equal(t, "division.bro.official.2017-beta", seasons[0].SeasonId)
	assert.Equal(t, true, seasons[25].IsCurrentSeason)
}
