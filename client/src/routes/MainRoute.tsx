import React, { useState, useEffect, useRef } from 'react';
import { Button, Select, MenuItem, Input, Container } from '@material-ui/core';
import type { PubgShard } from '../types/pubg';
import { useAsyncFn, useLocalStorage } from 'react-use';
import { fetchLifeTimeStats } from '../helper/stats';
import { LifeTimeStatsPanel } from '../components/LifeTimeStatsPanel';

export const MainRoute: React.FC = React.memo(() => {
  const [shard, setShard] = useState<PubgShard>('steam');
  const [username, setUsername] = useLocalStorage('queryUsername', '');

  const [state, fetch] = useAsyncFn(
    (renew: boolean) => {
      return fetchLifeTimeStats(shard, username, renew);
    },
    [shard, username]
  );

  useEffect(() => {
    if (username !== '') {
      fetch(false);
    }
  }, []);

  const isAutoLoadingRef = useRef(false);
  useEffect(() => {
    if (
      isAutoLoadingRef.current === false &&
      typeof state.value?.updatedAt === 'string' &&
      new Date().valueOf() - new Date(state.value?.updatedAt).valueOf() >
        1 * 24 * 60 * 60 * 1000
    ) {
      // 当前显示的统计数据时间距离当前时间已经超过1天
      // 则自动更新
      // 用于不手动更新时获取到过旧的数据
      fetch(true);
      isAutoLoadingRef.current = true; // 确保只自动刷新一次
      console.log('超过一天自动更新');
    }
  }, [state.value?.updatedAt]);

  return (
    <Container fixed>
      <div>Hello World</div>
      <Select
        value={shard}
        onChange={(e) => setShard(String(e.target.value) as PubgShard)}
      >
        <MenuItem value="steam">Steam</MenuItem>
      </Select>
      <Input
        placeholder="输入角色名..."
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <Button
        variant="contained"
        color="primary"
        disableElevation={true}
        disabled={state.loading}
        onClick={() => fetch(true)}
      >
        刷新请求
      </Button>

      {state.value && <LifeTimeStatsPanel stats={state.value} />}
    </Container>
  );
});
MainRoute.displayName = 'MainRoute';
