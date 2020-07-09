import React from 'react';
import { GameStats } from '../types/pubg';
import { getDurationTimeStr } from '../utils/time-helper';
import { Tip } from './Tip';
import { Grid, Typography, Divider, Box } from '@material-ui/core';

const StatsInfoGroup: React.FC<{
  title?: string;
}> = React.memo((props) => {
  return (
    <Box p={1}>
      <Typography>{props.title}</Typography>
      <Box marginTop={1} marginBottom={1}>
        <Divider />
      </Box>
      <Grid container spacing={3}>
        {props.children}
      </Grid>
    </Box>
  );
});
StatsInfoGroup.displayName = 'StatsInfoGroup';

const StatsInfoItem: React.FC = React.memo((props) => {
  return (
    <Grid item xs={3} container alignItems="center">
      {props.children}
    </Grid>
  );
});
StatsInfoItem.displayName = 'StatsInfoItem';

interface Props {
  data: GameStats;
}
export const StatsInfo: React.FC<Props> = React.memo((props) => {
  const { data } = props;
  return (
    <div>
      <StatsInfoGroup title="游戏记录">
        <StatsInfoItem>比赛次数: {data.roundsPlayed}</StatsInfoItem>

        <StatsInfoItem>游戏时间: {data.days}</StatsInfoItem>
        <StatsInfoItem>累积吃鸡: {data.wins}</StatsInfoItem>
        <StatsInfoItem>本日吃鸡: {data.dailyWins}</StatsInfoItem>
        <StatsInfoItem>本周吃鸡: {data.weeklyWins}</StatsInfoItem>
      </StatsInfoGroup>

      <StatsInfoGroup title="战斗">
        <StatsInfoItem>
          K/D: {Number((data.kills + data.assists) / data.losses).toFixed(2)}
        </StatsInfoItem>
        <StatsInfoItem>累积击杀: {data.kills}</StatsInfoItem>
        <StatsInfoItem>累积助攻: {data.assists}</StatsInfoItem>
        <StatsInfoItem>累积击倒: {data.dBNOs}</StatsInfoItem>

        <StatsInfoItem>本日击杀: {data.dailyKills}</StatsInfoItem>
        <StatsInfoItem>本周击杀: {data.weeklyKills}</StatsInfoItem>
        <StatsInfoItem>扶起队友: {data.revives}</StatsInfoItem>
        <StatsInfoItem>爆头击杀: {data.headshotKills}</StatsInfoItem>
        <StatsInfoItem>最远击杀: {data.longestKill}</StatsInfoItem>
        <StatsInfoItem>单场最高击杀: {data.roundMostKills}</StatsInfoItem>

        <StatsInfoItem>死亡次数: {data.losses}</StatsInfoItem>
        <StatsInfoItem>
          总伤害<Tip>该数据移除自身伤害</Tip>: {data.damageDealt}
        </StatsInfoItem>
      </StatsInfoGroup>

      <StatsInfoGroup title="历程">
        <StatsInfoItem>前10数: {data.top10s}</StatsInfoItem>
        <StatsInfoItem>载具击杀: {data.roadKills}</StatsInfoItem>
        <StatsInfoItem>使用能量饮料: {data.boosts}</StatsInfoItem>
        <StatsInfoItem>
          最长生存: {getDurationTimeStr(data.longestTimeSurvived)}
        </StatsInfoItem>
        <StatsInfoItem>治疗次数: {data.heals}</StatsInfoItem>
        <StatsInfoItem>载具摧毁: {data.vehicleDestroys}</StatsInfoItem>

        <StatsInfoItem>击杀队友: {data.teamKills}</StatsInfoItem>
        <StatsInfoItem>自杀次数: {data.suicides}</StatsInfoItem>
      </StatsInfoGroup>

      <StatsInfoGroup title="移动距离">
        <StatsInfoItem>步行距离: {data.walkDistance}</StatsInfoItem>
        <StatsInfoItem>驾驶距离: {data.rideDistance}</StatsInfoItem>
        <StatsInfoItem>游泳距离: {data.swimDistance}</StatsInfoItem>
      </StatsInfoGroup>
    </div>
  );
});
StatsInfo.displayName = 'StatsInfo';
