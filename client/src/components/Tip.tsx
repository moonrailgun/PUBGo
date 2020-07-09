import React from 'react';
import HelpOutline from '@material-ui/icons/HelpOutline';
import { Tooltip } from '@material-ui/core';

export const Tip: React.FC = React.memo((props) => {
  return (
    <Tooltip title={props.children} arrow placement="top">
      <HelpOutline fontSize="inherit">Arrow</HelpOutline>
    </Tooltip>
  );
});
Tip.displayName = 'Tip';
