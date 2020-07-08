import React from 'react';
import { Box, Typography, makeStyles } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  root: {
    overflow: 'auto',
    flex: 1,
  },
}));

interface Props {
  tabId: string;
  value: string;
}
export const TabPanel: React.FC<Props> = React.memo((props) => {
  const { children, tabId, value } = props;
  const classes = useStyles();

  return (
    <div
      className={classes.root}
      role="tabpanel"
      hidden={tabId !== value}
      id={`tabpanel-${value}`}
      aria-labelledby={`tab-${value}`}
    >
      {tabId === value && <Box p={3}>{children}</Box>}
    </div>
  );
});
TabPanel.displayName = 'TabPanel';
