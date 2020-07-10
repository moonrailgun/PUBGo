import { request } from './request';
import { PubgShard } from '../types/pubg';
import { quickParse } from '../utils/quickParse';

export interface MatchDetail {
  matchId: string;
  matchCreatedAt: string;
  duration: number;
  roster: MatchRoster[];
}

interface MatchRoster {
  id: string; // random uuid
  shardId: string;
  stats: {
    rank: number;
    teamId: number;
  };
  won: 'true' | 'false';
  participants: MatchParticipant[];
}

interface MatchParticipant {
  id: string; // random uuid
  actor: string;
  shardId: string;
}

/**
 * 获取战局信息
 */
export async function fetchMatchDetail(
  shard: PubgShard,
  matchId: string
): Promise<MatchDetail> {
  const { data } = await request.get(`/match/shard/${shard}/match/${matchId}`);
  data.roster = quickParse(data.roster);

  return data;
}
