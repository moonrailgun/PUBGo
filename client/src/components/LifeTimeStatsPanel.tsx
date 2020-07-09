import React, { useState } from 'react';
import { LifeTimeStats } from '../helper/stats';
import { getStandardTimeStr } from '../utils/time-helper';
import {
  Tab,
  Tabs,
  makeStyles,
  Typography,
  Grid,
  Switch,
} from '@material-ui/core';
import { TabPanel } from './TabPanel';
import { StatsInfo } from './StatsInfo';
import { useLocalStorage } from 'react-use';

const useStyles = makeStyles((theme) => ({
  root: {
    backgroundColor: theme.palette.background.paper,
  },
  updatedTime: {
    alignSelf: 'center',
  },
  tabsContainer: {
    flexGrow: 1,
    backgroundColor: theme.palette.background.paper,
    display: 'flex',
  },
  tabs: {
    borderRight: `1px solid ${theme.palette.divider}`,
  },
}));

interface Props {
  stats: LifeTimeStats;
}
export const LifeTimeStatsPanel: React.FC<Props> = React.memo((props) => {
  const { stats } = props;
  const [tabId, setTabId] = useLocalStorage('lifeTimeStatsTab', 'solo');
  const [isFpp, setIsFpp] = useLocalStorage('isFPP', false);
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container direction="row" justify="space-between">
        <Grid item className={classes.updatedTime}>
          <Typography>
            最后更新时间: {getStandardTimeStr(stats.updatedAt)}
          </Typography>
        </Grid>

        <Grid item>
          <Grid container alignItems="center" spacing={1}>
            <Grid item>TPP</Grid>
            <Grid item>
              <Switch
                checked={isFpp}
                onChange={(_, checked) => setIsFpp(checked)}
                name="FPP"
              />
            </Grid>
            <Grid item>FPP</Grid>
          </Grid>
        </Grid>
      </Grid>

      <div className={classes.tabsContainer}>
        <Tabs
          className={classes.tabs}
          orientation="vertical"
          variant="scrollable"
          value={tabId}
          onChange={(_, val) => setTabId(val)}
        >
          <Tab label="单排" value="solo" />
          <Tab label="双排" value="duo" />
          <Tab label="四排" value="squad" />
        </Tabs>

        <TabPanel tabId={tabId} value="solo">
          <StatsInfo
            data={
              isFpp ? stats.gameModeStats.soloFPP : stats.gameModeStats.solo
            }
          />
        </TabPanel>
        <TabPanel tabId={tabId} value="duo">
          <StatsInfo
            data={isFpp ? stats.gameModeStats.duoFPP : stats.gameModeStats.duo}
          />
        </TabPanel>
        <TabPanel tabId={tabId} value="squad">
          <StatsInfo
            data={
              isFpp ? stats.gameModeStats.squadFPP : stats.gameModeStats.squad
            }
          />
        </TabPanel>
      </div>
    </div>
  );
});
LifeTimeStatsPanel.displayName = 'LifeTimeStatsPanel';
