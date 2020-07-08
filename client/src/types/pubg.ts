export type PubgShard = 'steam';

export interface MatchSample {
  ID: string;
}

// 游戏统计数据
export interface GameStats {
  assists: number;
  boosts: number;
  dBNOs: number;
  dailyKills: number;
  damageDealt: number;
  days: number;
  dailyWins: number;
  headshotKills: number; // 爆头击杀数
  heals: number;
  killPoints: number;
  kills: number; // 击杀数
  longestKill: number; // 最远击杀
  longestTimeSurvived: number; // 最长生存时间
  losses: number;
  maxKillStreaks: number;
  mostSurvivalTime: number;
  rankPoints: number;
  rankPointsTitle: string;
  revives: number;
  rideDistance: number; // 载具移动距离
  roadKills: number;
  roundMostKills: number;
  roundsPlayed: number;
  suicides: number;
  swimDistance: number;
  teamKills: number;
  timeSurvived: number; // 累计生存时间
  top10s: number; // 前10数
  vehicleDestroys: number;
  walkDistance: number; // 累计步行距离
  weaponsAcquired: number;
  weeklyKills: number;
  weeklyWins: number;
  winPoints: number;
  wins: number; // 获胜次数
}
