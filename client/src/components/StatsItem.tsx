import React from 'react';
import { GameStats } from '../types/pubg';

interface Props {
  data: GameStats;
}
export const StatsItem: React.FC<Props> = React.memo((props) => {
  return <div>{JSON.stringify(props.data, null, 4)}</div>;
});
StatsItem.displayName = 'StatsItem';
