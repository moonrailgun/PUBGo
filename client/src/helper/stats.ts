import { PubgShard, MatchSample, GameStats } from '../types/pubg';
import { request } from './request';
import { quickParse } from '../utils/quickParse';

interface GameModeStats {
  solo: GameStats;
  soloFPP: GameStats;
  duo: GameStats;
  duoFPP: GameStats;
  squad: GameStats;
  squadFPP: GameStats;
}

export interface LifeTimeStats {
  accountId: string;
  updatedAt: string;
  gameModeStats: GameModeStats;
  id: number;
  matchesDuo: MatchSample[] | null;
  matchesDuoFPP: MatchSample[] | null;
  matchesSolo: MatchSample[] | null;
  matchesSoloFPP: MatchSample[] | null;
  matchesSquad: MatchSample[] | null;
  matchesSquadFPP: MatchSample[] | null;
}

/**
 * 获取生存时间统计
 */
export async function fetchLifeTimeStats(
  shard: PubgShard,
  username: string,
  renew: boolean = false
): Promise<LifeTimeStats> {
  let url = `/stats/shard/${shard}/name/${username}`;
  if (renew === true) {
    url += '/renew';
  }
  const { data } = await request.get(url);

  return {
    ...data,
    gameModeStats: quickParse(data.gameModeStats),
    matchesSolo: quickParse(data.matchesSolo),
    matchesSoloFPP: quickParse(data.matchesSoloFPP),
    matchesDuo: quickParse(data.matchesDuo),
    matchesDuoFPP: quickParse(data.matchesDuoFPP),
    matchesSquad: quickParse(data.matchesSquad),
    matchesSquadFPP: quickParse(data.matchesSquadFPP),
  };
}
