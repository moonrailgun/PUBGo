import React, { useCallback, useState } from 'react';
import { Button, Select, MenuItem, Input } from '@material-ui/core';
import { fetchPlayerInfo } from '../helper/player';
import type { PubgShard } from '../types/pubg';
import { useAsyncFn } from 'react-use';

export const MainRoute: React.FC = React.memo(() => {
  const [shard, setShard] = useState<PubgShard>('steam');
  const [username, setUsername] = useState('');

  const [state, fetch] = useAsyncFn(() => {
    return fetchPlayerInfo(shard, username);
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

      <div>{JSON.stringify(state.value)}</div>
    </div>
  );
});
MainRoute.displayName = 'MainRoute';
