import React from 'react';
import { MatchSample } from '../types/pubg';
import _isNil from 'lodash/isNil';
import { useAsyncFn, useAsync } from 'react-use';
import { fetchMatchDetail } from '../helper/match';

const MatchItem: React.FC<{
  matchId: string;
}> = React.memo((props) => {
  const { matchId } = props;

  const { value, loading, error } = useAsync(
    () => fetchMatchDetail('steam', matchId),
    [matchId]
  );

  if (loading) {
    return <div>加载中...</div>;
  }

  if (error) {
    return <div>出现错误</div>;
  }

  return <div>{JSON.stringify(value)}</div>;
});
MatchItem.displayName = 'MatchItem';

interface Props {
  matches: MatchSample[];
}
export const MatchList: React.FC<Props> = React.memo((props) => {
  const { matches } = props;

  if (_isNil(matches)) {
    return <div>暂无数据</div>;
  }

  return (
    <div>
      {matches.map((match) => {
        return <MatchItem matchId={match.id} />;
      })}
    </div>
  );
});
MatchList.displayName = 'MatchList';
