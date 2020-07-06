import type { PubgShard } from '../types/pubg';
import { request } from './request';

export async function fetchPlayerInfo(shard: PubgShard, username: string) {
  const { data } = await request.get(`/player/shard/${shard}/name/${username}`);

  return data;
}
