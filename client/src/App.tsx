import React from 'react';

import { NavBar } from './components/NavBar';
import CssBaseline from '@material-ui/core/CssBaseline';
import { MainRoute } from './routes/MainRoute';

export const App: React.FC = React.memo(() => {
  return (
    <div>
      {/* 基本CSS */}
      <CssBaseline />

      <NavBar />

      <MainRoute />
    </div>
  );
});
App.displayName = 'App';
