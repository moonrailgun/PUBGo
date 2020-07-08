import React, { useState } from 'react';
import { Button, Select, MenuItem, Input } from '@material-ui/core';
import type { PubgShard } from '../types/pubg';
import { useAsyncFn, useLocalStorage } from 'react-use';
import { fetchLifeTimeStats } from '../helper/stats';
import { LifeTimeStatsPanel } from '../components/LifeTimeStatsPanel';

export const MainRoute: React.FC = React.memo(() => {
  const [shard, setShard] = useState<PubgShard>('steam');
  const [username, setUsername] = useLocalStorage('queryUsername', '');

  const [state, fetch] = useAsyncFn(() => {
    return fetchLifeTimeStats(shard, username);
  }, [shard, username]);

  return (
    <div>
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
        onClick={fetch}
      >
        发送请求
      </Button>

      {state.value && <LifeTimeStatsPanel stats={state.value} />}
    </div>
  );
});
MainRoute.displayName = 'MainRoute';
