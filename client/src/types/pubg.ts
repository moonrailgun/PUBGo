export type PubgShard = 'steam';

export interface MatchSample {
  id: string;
}

// 基本统计数据
export interface BaseStats {
  dBNOs: number; // 击倒数
  assists: number; // 助攻数
  boosts: number;
  damageDealt: number;
  deathType: string;
  headshotKills: number; // 爆头击杀数
  heals: number;
  killPoints: number;
  kills: number; // 击杀数
  longestKill: number; // 最远击杀
  revives: number;
  rideDistance: number; // 载具移动距离
  roadKills: number;
  teamKills: number;
  timeSurvived: number; // 累计生存时间
  vehicleDestroys: number;
  walkDistance: number;
  weaponsAcquired: number;
  winPoints: number;
}

// 战局统计数据
export interface MatchStats extends BaseStats {
  killPlace: number;
  killPointsDelta: number;
  killStreaks: number;
  lastKillPoints: number;
  lastWinPoints: number;
  mostDamage: number;
  name: string;
  playerId: string;
  winPlace: number;
  winPointsDelta: number;
}

// 游戏统计数据
export interface GameStats extends BaseStats {
  dailyKills: number;
  days: number;
  dailyWins: number;
  longestTimeSurvived: number; // 最长生存时间
  losses: number;
  maxKillStreaks: number;
  mostSurvivalTime: number;
  rankPoints: number;
  rankPointsTitle: string;
  roundMostKills: number;
  roundsPlayed: number;
  suicides: number;
  swimDistance: number;
  top10s: number; // 前10数
  walkDistance: number; // 累计步行距离
  weeklyKills: number;
  weeklyWins: number;
  wins: number; // 获胜次数
}
