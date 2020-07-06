import type { PubgShard } from '../types/pubg';
import { request } from './request';

interface PlayerInfo {
  ID: string;
  Name: string;
  ShardID: string;
  CreatedAt: string;
  UpdatedAt: string;
  Matches: {
    ID: string;
  }[];
}

export async function fetchPlayerInfo(
  shard: PubgShard,
  username: string
): Promise<PlayerInfo> {
  const { data } = await request.get(`/player/shard/${shard}/name/${username}`);

  return data;
}
