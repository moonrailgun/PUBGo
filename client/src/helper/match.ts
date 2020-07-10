import { request } from './request';
import { PubgShard } from '../types/pubg';
import { quickParse } from '../utils/quickParse';

export interface MatchDetail {
  matchId: string;
  matchCreatedAt: string;
  duration: number;
  matchType: 'arcade' | 'custom' | 'event' | 'official' | 'training';
  mapName:
    | 'Baltic_Main'
    | 'Desert_Main'
    | 'DihorOtok_Main'
    | 'Erangel_Main'
    | 'Range_Main'
    | 'Savage_Main'
    | 'Summerland_Main';
  isCustomMatch: boolean;
  seasonState: 'closed' | 'prepare' | 'progress';
  roster: MatchRoster[];
  assets: MatchAssets[];
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

interface MatchAssets {
  id: string; // random uuid
  url: string;
  name: string; //telemetry
  description: string;
  createdAt: string;
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
  data.assets = quickParse(data.assets);

  return data;
}
