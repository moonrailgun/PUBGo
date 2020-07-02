package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testMatchStr = `
{
    "data": {
        "type": "match",
        "id": "186f4ee7-235c-4778-b581-8114821b8e05",
        "attributes": {
            "titleId": "bluehole-pubg",
            "mapName": "Summerland_Main",
            "matchType": "official",
            "seasonState": "progress",
            "stats": null,
            "duration": 1066,
            "gameMode": "squad",
            "shardId": "steam",
            "tags": null,
            "isCustomMatch": false,
            "createdAt": "2020-07-01T14:06:42Z"
        },
        "relationships": {
            "assets": {
                "data": [
                    {
                        "type": "asset",
                        "id": "d7b3863a-bba6-11ea-9fb1-520fabc4fe4b"
                    }
                ]
            },
            "rosters": {
                "data": [
                    {
                        "type": "roster",
                        "id": "e4ad0cc5-d9e3-4b7b-86b0-b9ef360df8c4"
                    },
                    {
                        "type": "roster",
                        "id": "09d1c9cd-eadb-41d1-b383-46157ad2653b"
                    },
                    {
                        "type": "roster",
                        "id": "4f34d068-212b-4863-980d-ef34e4f367fd"
                    },
                    {
                        "type": "roster",
                        "id": "a40d1a6b-1cb0-4d75-940f-3c258b126b1c"
                    },
                    {
                        "type": "roster",
                        "id": "123661df-6fb0-476c-9eaf-3e8d2c45762a"
                    },
                    {
                        "type": "roster",
                        "id": "2e1e6e22-3d5e-4016-a209-7efd6a3bbcea"
                    },
                    {
                        "type": "roster",
                        "id": "2c81a14d-09ff-4319-b4ab-4e0cb998fea6"
                    },
                    {
                        "type": "roster",
                        "id": "2146698a-cf49-4be6-9ed7-9a238cad78ba"
                    },
                    {
                        "type": "roster",
                        "id": "670775c1-123c-4b84-b34b-3ad77178be98"
                    }
                ]
            }
        },
        "links": {
            "self": "https://api.playbattlegrounds.com/shards/steam/matches/186f4ee7-235c-4778-b581-8114821b8e05",
            "schema": ""
        }
    },
    "included": [
        {
            "type": "participant",
            "id": "9d8bf1f9-dee5-4f88-aa92-32534dbd118d",
            "attributes": {
                "stats": {
                    "DBNOs": 1,
                    "assists": 0,
                    "boosts": 5,
                    "damageDealt": 370.29413,
                    "deathType": "byplayer",
                    "headshotKills": 1,
                    "heals": 5,
                    "killPlace": 10,
                    "killStreaks": 0,
                    "kills": 2,
                    "longestKill": 27.458864,
                    "name": "Monkey_c9",
                    "playerId": "account.40b38bcb5f974cdf9c2f0e35e665c926",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 653.241,
                    "vehicleDestroys": 0,
                    "walkDistance": 1168.8845,
                    "weaponsAcquired": 6,
                    "winPlace": 6
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "14626f08-0f79-4abd-bba6-24217ca4ae0b",
            "attributes": {
                "shardId": "steam",
                "stats": {
                    "DBNOs": 2,
                    "assists": 0,
                    "boosts": 6,
                    "damageDealt": 226.16246,
                    "deathType": "byplayer",
                    "headshotKills": 1,
                    "heals": 17,
                    "killPlace": 7,
                    "killStreaks": 0,
                    "kills": 3,
                    "longestKill": 31.848413,
                    "name": "jiangxing121",
                    "playerId": "account.98f8f212ca0f4e40a33cf48446f0bdb0",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 1053.921,
                    "vehicleDestroys": 0,
                    "walkDistance": 1659.6019,
                    "weaponsAcquired": 6,
                    "winPlace": 3
                },
                "actor": ""
            }
        },
        {
            "type": "roster",
            "id": "09d1c9cd-eadb-41d1-b383-46157ad2653b",
            "attributes": {
                "stats": {
                    "rank": 5,
                    "teamId": 6
                },
                "won": "false",
                "shardId": "steam"
            },
            "relationships": {
                "team": {
                    "data": null
                },
                "participants": {
                    "data": [
                        {
                            "type": "participant",
                            "id": "fd79f1cb-95d3-488f-92df-115914e9cb6f"
                        },
                        {
                            "type": "participant",
                            "id": "f0687147-9009-4c0d-8626-382c7d1b1aa9"
                        },
                        {
                            "type": "participant",
                            "id": "9935d2d6-39b8-4542-95b2-604ffe71b83b"
                        },
                        {
                            "type": "participant",
                            "id": "2c9f15cc-7e2c-4da7-93d7-c7944d4bfcb6"
                        }
                    ]
                }
            }
        },
        {
            "type": "roster",
            "id": "4f34d068-212b-4863-980d-ef34e4f367fd",
            "attributes": {
                "won": "true",
                "shardId": "steam",
                "stats": {
                    "rank": 1,
                    "teamId": 8
                }
            },
            "relationships": {
                "participants": {
                    "data": [
                        {
                            "type": "participant",
                            "id": "16d449c2-99be-49a6-8aef-c9b0c153b748"
                        },
                        {
                            "type": "participant",
                            "id": "8e8f7684-8cef-48e2-86ef-909cc35ddd3e"
                        },
                        {
                            "type": "participant",
                            "id": "d8821bab-f248-4c0e-9a6c-d80a5619aeb0"
                        },
                        {
                            "type": "participant",
                            "id": "6a38aa91-2d92-4815-a925-2080a8caa15e"
                        }
                    ]
                },
                "team": {
                    "data": null
                }
            }
        },
        {
            "type": "participant",
            "id": "2c9f15cc-7e2c-4da7-93d7-c7944d4bfcb6",
            "attributes": {
                "stats": {
                    "DBNOs": 5,
                    "assists": 1,
                    "boosts": 5,
                    "damageDealt": 676.3915,
                    "deathType": "byplayer",
                    "headshotKills": 1,
                    "heals": 2,
                    "killPlace": 2,
                    "killStreaks": 0,
                    "kills": 8,
                    "longestKill": 152.41101,
                    "name": "EDG_CIearlove777",
                    "playerId": "account.1321395a9f9e4c5ea006f983413c2894",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 770.471,
                    "vehicleDestroys": 0,
                    "walkDistance": 1719.9644,
                    "weaponsAcquired": 5,
                    "winPlace": 5
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "8e8f7684-8cef-48e2-86ef-909cc35ddd3e",
            "attributes": {
                "actor": "",
                "shardId": "steam",
                "stats": {
                    "DBNOs": 0,
                    "assists": 0,
                    "boosts": 4,
                    "damageDealt": 104.46381,
                    "deathType": "alive",
                    "headshotKills": 1,
                    "heals": 0,
                    "killPlace": 14,
                    "killStreaks": 0,
                    "kills": 1,
                    "longestKill": 38.742588,
                    "name": "Morning_Star_ll",
                    "playerId": "account.40aeeb9846774412abf9a320d575fbe1",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 1066.114,
                    "vehicleDestroys": 0,
                    "walkDistance": 923.4704,
                    "weaponsAcquired": 6,
                    "winPlace": 1
                }
            }
        },
        {
            "type": "participant",
            "id": "ff405b35-6849-464e-8518-7522da231bd4",
            "attributes": {
                "stats": {
                    "DBNOs": 1,
                    "assists": 0,
                    "boosts": 1,
                    "damageDealt": 381.72467,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 7,
                    "killPlace": 5,
                    "killStreaks": 0,
                    "kills": 3,
                    "longestKill": 140.6486,
                    "name": "zhenping0377",
                    "playerId": "account.7560310e55894fd6b384aa001fcc2749",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 1057.972,
                    "vehicleDestroys": 0,
                    "walkDistance": 1807.9204,
                    "weaponsAcquired": 4,
                    "winPlace": 2
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "ae09926e-8e54-45af-a78f-ebd8c037ab2f",
            "attributes": {
                "shardId": "steam",
                "stats": {
                    "DBNOs": 0,
                    "assists": 1,
                    "boosts": 1,
                    "damageDealt": 57.809998,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 26,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "liangwei198",
                    "playerId": "account.5489b15d78a5478285a4251a6758bf0b",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 374.938,
                    "vehicleDestroys": 0,
                    "walkDistance": 877.76904,
                    "weaponsAcquired": 2,
                    "winPlace": 2
                },
                "actor": ""
            }
        },
        {
            "type": "participant",
            "id": "c0b8c29b-262f-4193-ad2e-b95b7957c0b0",
            "attributes": {
                "stats": {
                    "DBNOs": 0,
                    "assists": 0,
                    "boosts": 2,
                    "damageDealt": 0,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 32,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "Mxx_Deng_Bao",
                    "playerId": "account.e8eaacef22ef421d98bda4593c5f144c",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 588.183,
                    "vehicleDestroys": 0,
                    "walkDistance": 994.9897,
                    "weaponsAcquired": 4,
                    "winPlace": 6
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "a2114615-b779-4485-9553-7db612906be3",
            "attributes": {
                "stats": {
                    "DBNOs": 0,
                    "assists": 0,
                    "boosts": 4,
                    "damageDealt": 88.79276,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 29,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "Alone_Ki",
                    "playerId": "account.4b293bf3cdad4f88ae0031fc7df501a0",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 700.743,
                    "vehicleDestroys": 0,
                    "walkDistance": 1312.1215,
                    "weaponsAcquired": 6,
                    "winPlace": 6
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "37fb2728-e45d-48b1-a050-e40724a7e6c5",
            "attributes": {
                "actor": "",
                "shardId": "steam",
                "stats": {
                    "DBNOs": 3,
                    "assists": 1,
                    "boosts": 2,
                    "damageDealt": 231.06001,
                    "deathType": "byplayer",
                    "headshotKills": 1,
                    "heals": 1,
                    "killPlace": 11,
                    "killStreaks": 0,
                    "kills": 2,
                    "longestKill": 38.03994,
                    "name": "AqnaLotus",
                    "playerId": "account.5c3d60147eb549d18a0984a67af3f58b",
                    "revives": 1,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 571.488,
                    "vehicleDestroys": 0,
                    "walkDistance": 915.54474,
                    "weaponsAcquired": 3,
                    "winPlace": 6
                }
            }
        },
        {
            "type": "roster",
            "id": "670775c1-123c-4b84-b34b-3ad77178be98",
            "attributes": {
                "won": "false",
                "shardId": "steam",
                "stats": {
                    "rank": 4,
                    "teamId": 2
                }
            },
            "relationships": {
                "team": {
                    "data": null
                },
                "participants": {
                    "data": [
                        {
                            "type": "participant",
                            "id": "dbe0c05e-d4a0-4005-9c1f-10d0d402fff5"
                        },
                        {
                            "type": "participant",
                            "id": "37fb2728-e45d-48b1-a050-e40724a7e6c5"
                        },
                        {
                            "type": "participant",
                            "id": "381037e0-6c6f-405d-abeb-f67018ce76cd"
                        },
                        {
                            "type": "participant",
                            "id": "fb50c981-bbc6-4b3f-8c94-4a8d69bbcd31"
                        }
                    ]
                }
            }
        },
        {
            "type": "participant",
            "id": "43b9c76e-d96c-4323-a03a-a0ae77091829",
            "attributes": {
                "stats": {
                    "DBNOs": 0,
                    "assists": 1,
                    "boosts": 1,
                    "damageDealt": 27.060001,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 48,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "QinDaShuaiGe98-",
                    "playerId": "account.75c6f6c7d15e480dbaea1f4e96de4926",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 237.918,
                    "vehicleDestroys": 0,
                    "walkDistance": 557.27185,
                    "weaponsAcquired": 3,
                    "winPlace": 14
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "16d449c2-99be-49a6-8aef-c9b0c153b748",
            "attributes": {
                "stats": {
                    "DBNOs": 6,
                    "assists": 2,
                    "boosts": 3,
                    "damageDealt": 906.76733,
                    "deathType": "alive",
                    "headshotKills": 4,
                    "heals": 3,
                    "killPlace": 1,
                    "killStreaks": 0,
                    "kills": 8,
                    "longestKill": 80.768845,
                    "name": "bigJIJIXIA",
                    "playerId": "account.6fce248ccead4a9ba7b176e5f531cbe6",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 1066.114,
                    "vehicleDestroys": 0,
                    "walkDistance": 1677.7067,
                    "weaponsAcquired": 4,
                    "winPlace": 1
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "roster",
            "id": "2146698a-cf49-4be6-9ed7-9a238cad78ba",
            "attributes": {
                "won": "false",
                "shardId": "steam",
                "stats": {
                    "rank": 3,
                    "teamId": 4
                }
            },
            "relationships": {
                "team": {
                    "data": null
                },
                "participants": {
                    "data": [
                        {
                            "type": "participant",
                            "id": "96017872-8bee-49cd-9727-5f79a9e1952c"
                        },
                        {
                            "type": "participant",
                            "id": "46b89393-1cbf-4084-abbd-000e5f2b6ef5"
                        },
                        {
                            "type": "participant",
                            "id": "9d8bf1f9-dee5-4f88-aa92-32534dbd118d"
                        },
                        {
                            "type": "participant",
                            "id": "14626f08-0f79-4abd-bba6-24217ca4ae0b"
                        }
                    ]
                }
            }
        },
        {
            "type": "asset",
            "id": "d7b3863a-bba6-11ea-9fb1-520fabc4fe4b",
            "attributes": {
                "URL": "https://telemetry-cdn.playbattlegrounds.com/bluehole-pubg/steam/2020/07/01/14/26/d7b3863a-bba6-11ea-9fb1-520fabc4fe4b-telemetry.json",
                "name": "telemetry",
                "description": "",
                "createdAt": "2020-07-01T14:26:24Z"
            }
        },
        {
            "type": "participant",
            "id": "d8821bab-f248-4c0e-9a6c-d80a5619aeb0",
            "attributes": {
                "stats": {
                    "DBNOs": 0,
                    "assists": 2,
                    "boosts": 4,
                    "damageDealt": 408.8842,
                    "deathType": "alive",
                    "headshotKills": 0,
                    "heals": 3,
                    "killPlace": 9,
                    "killStreaks": 0,
                    "kills": 2,
                    "longestKill": 97.90725,
                    "name": "zsy54007",
                    "playerId": "account.6f1cc10bc82147feb63cc1ee04c775c4",
                    "revives": 1,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 1066.114,
                    "vehicleDestroys": 0,
                    "walkDistance": 1644.2087,
                    "weaponsAcquired": 4,
                    "winPlace": 1
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "769867a4-6368-4824-b6eb-51fe15a9ad51",
            "attributes": {
                "actor": "",
                "shardId": "steam",
                "stats": {
                    "DBNOs": 0,
                    "assists": 0,
                    "boosts": 0,
                    "damageDealt": 82.6,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 62,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "Snister-blade",
                    "playerId": "account.29014331fe5e41779902cb760b25f0df",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 191.448,
                    "vehicleDestroys": 0,
                    "walkDistance": 224.08295,
                    "weaponsAcquired": 2,
                    "winPlace": 17
                }
            }
        },
        {
            "type": "participant",
            "id": "381037e0-6c6f-405d-abeb-f67018ce76cd",
            "attributes": {
                "stats": {
                    "DBNOs": 0,
                    "assists": 0,
                    "boosts": 0,
                    "damageDealt": 0,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 33,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "kgsgma",
                    "playerId": "account.9c1b77a043b445af8edc1f0e593dc40b",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 69.991,
                    "vehicleDestroys": 0,
                    "walkDistance": 4.2882137,
                    "weaponsAcquired": 0,
                    "winPlace": 6
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "roster",
            "id": "2e1e6e22-3d5e-4016-a209-7efd6a3bbcea",
            "attributes": {
                "shardId": "steam",
                "stats": {
                    "rank": 6,
                    "teamId": 9
                },
                "won": "false"
            },
            "relationships": {
                "participants": {
                    "data": [
                        {
                            "type": "participant",
                            "id": "c0b8c29b-262f-4193-ad2e-b95b7957c0b0"
                        },
                        {
                            "type": "participant",
                            "id": "a2114615-b779-4485-9553-7db612906be3"
                        },
                        {
                            "type": "participant",
                            "id": "04dfd9a2-f111-42f0-920c-1e080091be88"
                        },
                        {
                            "type": "participant",
                            "id": "dd4edca9-0d4a-4ed8-ac45-2be94cc2aee6"
                        }
                    ]
                },
                "team": {
                    "data": null
                }
            }
        },
        {
            "type": "participant",
            "id": "f0687147-9009-4c0d-8626-382c7d1b1aa9",
            "attributes": {
                "stats": {
                    "DBNOs": 5,
                    "assists": 1,
                    "boosts": 5,
                    "damageDealt": 556.51526,
                    "deathType": "byplayer",
                    "headshotKills": 1,
                    "heals": 1,
                    "killPlace": 3,
                    "killStreaks": 0,
                    "kills": 5,
                    "longestKill": 202.43933,
                    "name": "PtwoY",
                    "playerId": "account.ab4c06e9c951484a9a8cb5a57b8d8239",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 720.771,
                    "vehicleDestroys": 0,
                    "walkDistance": 2040.4606,
                    "weaponsAcquired": 5,
                    "winPlace": 5
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "a553caf9-2430-4d29-a709-90ea166c53d5",
            "attributes": {
                "stats": {
                    "DBNOs": 1,
                    "assists": 0,
                    "boosts": 1,
                    "damageDealt": 45.116867,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 63,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "you--summer",
                    "playerId": "account.78d73d7855814dac893ee330c67e4980",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 191.443,
                    "vehicleDestroys": 0,
                    "walkDistance": 179.77893,
                    "weaponsAcquired": 2,
                    "winPlace": 17
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "04dfd9a2-f111-42f0-920c-1e080091be88",
            "attributes": {
                "stats": {
                    "DBNOs": 0,
                    "assists": 0,
                    "boosts": 4,
                    "damageDealt": 34.8975,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 1,
                    "killPlace": 31,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "JunJun_Girl",
                    "playerId": "account.2bfafd28ad6744a6a6376b0dbe6aef15",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 699.229,
                    "vehicleDestroys": 0,
                    "walkDistance": 1577.3347,
                    "weaponsAcquired": 5,
                    "winPlace": 6
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "dd4edca9-0d4a-4ed8-ac45-2be94cc2aee6",
            "attributes": {
                "actor": "",
                "shardId": "steam",
                "stats": {
                    "DBNOs": 0,
                    "assists": 0,
                    "boosts": 4,
                    "damageDealt": 0,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 1,
                    "killPlace": 30,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "xttls",
                    "playerId": "account.f600fac9157848c0a4cf08ee02f5e719",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 700.731,
                    "vehicleDestroys": 0,
                    "walkDistance": 1566.1306,
                    "weaponsAcquired": 5,
                    "winPlace": 6
                }
            }
        },
        {
            "type": "participant",
            "id": "46b89393-1cbf-4084-abbd-000e5f2b6ef5",
            "attributes": {
                "shardId": "steam",
                "stats": {
                    "DBNOs": 2,
                    "assists": 0,
                    "boosts": 0,
                    "damageDealt": 426.3638,
                    "deathType": "byplayer",
                    "headshotKills": 2,
                    "heals": 0,
                    "killPlace": 4,
                    "killStreaks": 0,
                    "kills": 4,
                    "longestKill": 83.512596,
                    "name": "womendouyiyang11",
                    "playerId": "account.4214fa3d00f04ee4830038b363413ee4",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 486.709,
                    "vehicleDestroys": 0,
                    "walkDistance": 867.747,
                    "weaponsAcquired": 4,
                    "winPlace": 7
                },
                "actor": ""
            }
        },
        {
            "type": "roster",
            "id": "2c81a14d-09ff-4319-b4ab-4e0cb998fea6",
            "attributes": {
                "stats": {
                    "rank": 17,
                    "teamId": 5
                },
                "won": "false",
                "shardId": "steam"
            },
            "relationships": {
                "team": {
                    "data": null
                },
                "participants": {
                    "data": [
                        {
                            "type": "participant",
                            "id": "769867a4-6368-4824-b6eb-51fe15a9ad51"
                        },
                        {
                            "type": "participant",
                            "id": "a553caf9-2430-4d29-a709-90ea166c53d5"
                        },
                        {
                            "type": "participant",
                            "id": "e333de8d-b23b-4d62-824c-bf7c86748ff6"
                        }
                    ]
                }
            }
        },
        {
            "type": "participant",
            "id": "e0a1af9c-5cc3-4137-8417-7eda98e6ad2b",
            "attributes": {
                "stats": {
                    "DBNOs": 1,
                    "assists": 1,
                    "boosts": 0,
                    "damageDealt": 42.190002,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 1,
                    "killPlace": 25,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "JohnDaisy",
                    "playerId": "account.6127f7b215b14462a1ee676250b5445b",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 385.242,
                    "vehicleDestroys": 0,
                    "walkDistance": 748.29736,
                    "weaponsAcquired": 4,
                    "winPlace": 2
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "15745dca-ee5d-471d-a90f-3ef3b5d15213",
            "attributes": {
                "stats": {
                    "DBNOs": 1,
                    "assists": 0,
                    "boosts": 3,
                    "damageDealt": 254.12943,
                    "deathType": "byzone",
                    "headshotKills": 0,
                    "heals": 3,
                    "killPlace": 24,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "iAm-Cute",
                    "playerId": "account.e9d5adc7f4ea48248fe47f9c8cf38771",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 679.851,
                    "vehicleDestroys": 0,
                    "walkDistance": 1323.3121,
                    "weaponsAcquired": 4,
                    "winPlace": 2
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "fd79f1cb-95d3-488f-92df-115914e9cb6f",
            "attributes": {
                "stats": {
                    "DBNOs": 1,
                    "assists": 1,
                    "boosts": 5,
                    "damageDealt": 17.687878,
                    "deathType": "byplayer",
                    "headshotKills": 1,
                    "heals": 2,
                    "killPlace": 15,
                    "killStreaks": 0,
                    "kills": 1,
                    "longestKill": 44.55987,
                    "name": "K7-II",
                    "playerId": "account.8e4f9ed147964b6db3be82d27c85ffbc",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 741.331,
                    "vehicleDestroys": 0,
                    "walkDistance": 1824.5564,
                    "weaponsAcquired": 6,
                    "winPlace": 5
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "af5fd174-d59a-49ed-b800-4ce72d9aa2a4",
            "attributes": {
                "stats": {
                    "DBNOs": 0,
                    "assists": 0,
                    "boosts": 0,
                    "damageDealt": 0,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 64,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "Black-Stone-",
                    "playerId": "account.457aa9e0f2484394b1aa15b4537a8924",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 112.886,
                    "vehicleDestroys": 0,
                    "walkDistance": 0,
                    "weaponsAcquired": 0,
                    "winPlace": 17
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "dbe0c05e-d4a0-4005-9c1f-10d0d402fff5",
            "attributes": {
                "stats": {
                    "DBNOs": 2,
                    "assists": 1,
                    "boosts": 3,
                    "damageDealt": 461.72476,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 3,
                    "killPlace": 8,
                    "killStreaks": 0,
                    "kills": 3,
                    "longestKill": 37.039932,
                    "name": "2222222-----",
                    "playerId": "account.88b60c0cfa494d35a5f9902cd1e05680",
                    "revives": 1,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 572.089,
                    "vehicleDestroys": 0,
                    "walkDistance": 1035.6476,
                    "weaponsAcquired": 5,
                    "winPlace": 6
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "fb50c981-bbc6-4b3f-8c94-4a8d69bbcd31",
            "attributes": {
                "stats": {
                    "DBNOs": 0,
                    "assists": 1,
                    "boosts": 4,
                    "damageDealt": 22.140001,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 1,
                    "killPlace": 27,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "Evenmo",
                    "playerId": "account.b764a90221474565b308cfd7ec2572b7",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 1011.418,
                    "vehicleDestroys": 0,
                    "walkDistance": 2610.0417,
                    "weaponsAcquired": 2,
                    "winPlace": 4
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "roster",
            "id": "123661df-6fb0-476c-9eaf-3e8d2c45762a",
            "attributes": {
                "shardId": "steam",
                "stats": {
                    "rank": 17,
                    "teamId": 10
                },
                "won": "false"
            },
            "relationships": {
                "team": {
                    "data": null
                },
                "participants": {
                    "data": [
                        {
                            "type": "participant",
                            "id": "7d46713d-7fb8-4ea7-b148-d50ef6870b11"
                        },
                        {
                            "type": "participant",
                            "id": "af5fd174-d59a-49ed-b800-4ce72d9aa2a4"
                        }
                    ]
                }
            }
        },
        {
            "type": "participant",
            "id": "1044261e-6670-4966-972d-4a04e0b3a306",
            "attributes": {
                "stats": {
                    "DBNOs": 2,
                    "assists": 1,
                    "boosts": 0,
                    "damageDealt": 205.76222,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 22,
                    "killStreaks": 0,
                    "kills": 1,
                    "longestKill": 22.194773,
                    "name": "robot_007_007",
                    "playerId": "account.5c1b9ed06f5640c3980c5e38dd084933",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 259.281,
                    "vehicleDestroys": 0,
                    "walkDistance": 552.4957,
                    "weaponsAcquired": 2,
                    "winPlace": 13
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "27610094-983a-4a22-9b15-1d5f45784f59",
            "attributes": {
                "stats": {
                    "DBNOs": 1,
                    "assists": 0,
                    "boosts": 0,
                    "damageDealt": 75,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 21,
                    "killStreaks": 0,
                    "kills": 1,
                    "longestKill": 20.481298,
                    "name": "Bear188ZP",
                    "playerId": "account.dac2be4f6ef0493ab783745f462e6441",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 278.504,
                    "vehicleDestroys": 0,
                    "walkDistance": 653.24994,
                    "weaponsAcquired": 3,
                    "winPlace": 13
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "7d46713d-7fb8-4ea7-b148-d50ef6870b11",
            "attributes": {
                "stats": {
                    "DBNOs": 0,
                    "assists": 0,
                    "boosts": 0,
                    "damageDealt": 0,
                    "deathType": "byzone",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 60,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "YUSHENG_YUNI",
                    "playerId": "account.9ad1388cd37f40fb91e905af37569be1",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 416.475,
                    "vehicleDestroys": 0,
                    "walkDistance": 0,
                    "weaponsAcquired": 0,
                    "winPlace": 17
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "96017872-8bee-49cd-9727-5f79a9e1952c",
            "attributes": {
                "stats": {
                    "DBNOs": 1,
                    "assists": 0,
                    "boosts": 3,
                    "damageDealt": 175.60138,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 1,
                    "killPlace": 16,
                    "killStreaks": 0,
                    "kills": 1,
                    "longestKill": 6.6653614,
                    "name": "chengbeiopanan",
                    "playerId": "account.f44c3e37f8c1422684482b0572f01632",
                    "revives": 1,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 619.632,
                    "vehicleDestroys": 0,
                    "walkDistance": 1033.3352,
                    "weaponsAcquired": 6,
                    "winPlace": 6
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "roster",
            "id": "a40d1a6b-1cb0-4d75-940f-3c258b126b1c",
            "attributes": {
                "stats": {
                    "rank": 2,
                    "teamId": 7
                },
                "won": "false",
                "shardId": "steam"
            },
            "relationships": {
                "team": {
                    "data": null
                },
                "participants": {
                    "data": [
                        {
                            "type": "participant",
                            "id": "e0a1af9c-5cc3-4137-8417-7eda98e6ad2b"
                        },
                        {
                            "type": "participant",
                            "id": "ff405b35-6849-464e-8518-7522da231bd4"
                        },
                        {
                            "type": "participant",
                            "id": "15745dca-ee5d-471d-a90f-3ef3b5d15213"
                        },
                        {
                            "type": "participant",
                            "id": "ae09926e-8e54-45af-a78f-ebd8c037ab2f"
                        }
                    ]
                }
            }
        },
        {
            "type": "participant",
            "id": "9935d2d6-39b8-4542-95b2-604ffe71b83b",
            "attributes": {
                "stats": {
                    "DBNOs": 1,
                    "assists": 2,
                    "boosts": 3,
                    "damageDealt": 98.92317,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 28,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "TaoYing_1",
                    "playerId": "account.c41bffb890374d9fb57b3031001fb3b2",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 770.46,
                    "vehicleDestroys": 0,
                    "walkDistance": 2063.5242,
                    "weaponsAcquired": 3,
                    "winPlace": 5
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "participant",
            "id": "6a38aa91-2d92-4815-a925-2080a8caa15e",
            "attributes": {
                "stats": {
                    "DBNOs": 3,
                    "assists": 1,
                    "boosts": 3,
                    "damageDealt": 372.64417,
                    "deathType": "byzone",
                    "headshotKills": 0,
                    "heals": 2,
                    "killPlace": 6,
                    "killStreaks": 0,
                    "kills": 3,
                    "longestKill": 129.8761,
                    "name": "My_rookie",
                    "playerId": "account.d9b455971f1f4c2da24e9385a7344e94",
                    "revives": 1,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 846.941,
                    "vehicleDestroys": 0,
                    "walkDistance": 1254.064,
                    "weaponsAcquired": 2,
                    "winPlace": 2
                },
                "actor": "",
                "shardId": "steam"
            }
        },
        {
            "type": "roster",
            "id": "e4ad0cc5-d9e3-4b7b-86b0-b9ef360df8c4",
            "attributes": {
                "shardId": "steam",
                "stats": {
                    "rank": 13,
                    "teamId": 3
                },
                "won": "false"
            },
            "relationships": {
                "team": {
                    "data": null
                },
                "participants": {
                    "data": [
                        {
                            "type": "participant",
                            "id": "94bdfa18-c0da-4bf1-bb0a-ac772962f84e"
                        },
                        {
                            "type": "participant",
                            "id": "1044261e-6670-4966-972d-4a04e0b3a306"
                        },
                        {
                            "type": "participant",
                            "id": "43b9c76e-d96c-4323-a03a-a0ae77091829"
                        },
                        {
                            "type": "participant",
                            "id": "27610094-983a-4a22-9b15-1d5f45784f59"
                        }
                    ]
                }
            }
        },
        {
            "type": "participant",
            "id": "94bdfa18-c0da-4bf1-bb0a-ac772962f84e",
            "attributes": {
                "shardId": "steam",
                "stats": {
                    "DBNOs": 3,
                    "assists": 0,
                    "boosts": 0,
                    "damageDealt": 291.77777,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 13,
                    "killStreaks": 0,
                    "kills": 2,
                    "longestKill": 91.08525,
                    "name": "masajin",
                    "playerId": "account.5438ea9a85d44b5abd1c1b63223120a4",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 335.741,
                    "vehicleDestroys": 0,
                    "walkDistance": 716.6611,
                    "weaponsAcquired": 2,
                    "winPlace": 13
                },
                "actor": ""
            }
        },
        {
            "type": "participant",
            "id": "e333de8d-b23b-4d62-824c-bf7c86748ff6",
            "attributes": {
                "stats": {
                    "DBNOs": 0,
                    "assists": 0,
                    "boosts": 0,
                    "damageDealt": 107.40001,
                    "deathType": "byplayer",
                    "headshotKills": 0,
                    "heals": 0,
                    "killPlace": 61,
                    "killStreaks": 0,
                    "kills": 0,
                    "longestKill": 0,
                    "name": "41sbJJ",
                    "playerId": "account.b765f77509da42a3857f7264d3d4ff7b",
                    "revives": 0,
                    "rideDistance": 0,
                    "roadKills": 0,
                    "swimDistance": 0,
                    "teamKills": 0,
                    "timeSurvived": 191.458,
                    "vehicleDestroys": 0,
                    "walkDistance": 292.02008,
                    "weaponsAcquired": 2,
                    "winPlace": 17
                },
                "actor": "",
                "shardId": "steam"
            }
        }
    ],
    "links": {
        "self": "https://api-origin.playbattlegrounds.com/shards/steam/matches/186f4ee7-235c-4778-b581-8114821b8e05"
    },
    "meta": {
    }
}
`

func TestParseMatch(t *testing.T) {
	match, err := ParseMatch(testMatchStr); if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "186f4ee7-235c-4778-b581-8114821b8e05", match.ID)
	assert.Equal(t, 1066, match.Duration)
	assert.Equal(t, "squad", match.GameMode)
	assert.Equal(t, "steam", match.ShardID)
	assert.Equal(t, "bluehole-pubg", match.TitleID)
	assert.Equal(t, "e4ad0cc5-d9e3-4b7b-86b0-b9ef360df8c4", match.Rosters[0].ID)
	assert.Equal(t, "false", match.Rosters[0].Won)
	assert.Len(t, match.Rosters, 9)
}