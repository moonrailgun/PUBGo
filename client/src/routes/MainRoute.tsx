import React, { useCallback, useState } from 'react';
import { Button, Select, MenuItem, Input } from '@material-ui/core';
import { fetchPlayerInfo } from '../helper/player';
import type { PubgShard } from '../types/pubg';

export const MainRoute: React.FC = React.memo(() => {
  const [shard, setShard] = useState<PubgShard>('steam');
  const [username, setUsername] = useState('');

  const [res, setRes] = useState('');

  const handleClick = useCallback(() => {
    fetchPlayerInfo(shard, username).then((data) => setRes(data));
  }, [shard, username, setRes]);

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
        onClick={handleClick}
      >
        发送请求
      </Button>

      <div>{JSON.stringify(res)}</div>
    </div>
  );
});
MainRoute.displayName = 'MainRoute';
