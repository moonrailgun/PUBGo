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
import { StatsItem } from './StatsItem';

const useStyles = makeStyles((theme) => ({
  tabsContainer: {
    flexGrow: 1,
    backgroundColor: theme.palette.background.paper,
    display: 'flex',
    height: 224,
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
  const [tabId, setTabId] = useState('solo');
  const [isFpp, setIsFpp] = useState(false);
  const classes = useStyles();

  return (
    <div>
      <Typography>
        最后更新时间: {getStandardTimeStr(stats.updatedAt)}
      </Typography>

      <Grid component="label" container alignItems="center" spacing={1}>
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
          <StatsItem
            data={
              isFpp ? stats.gameModeStats.soloFPP : stats.gameModeStats.solo
            }
          />
        </TabPanel>
        <TabPanel tabId={tabId} value="duo">
          <StatsItem
            data={isFpp ? stats.gameModeStats.duoFPP : stats.gameModeStats.duo}
          />
        </TabPanel>
        <TabPanel tabId={tabId} value="squad">
          <StatsItem
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
