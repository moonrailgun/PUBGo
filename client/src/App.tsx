import React from 'react';
import { NavBar } from './components/NavBar';
import CssBaseline from '@material-ui/core/CssBaseline';
import { MainRoute } from './routes/MainRoute';
import { ThemeProvider } from '@material-ui/core';
import { theme } from './theme';

const Provider: React.FC = React.memo((props) => {
  return <ThemeProvider theme={theme}>{props.children}</ThemeProvider>;
});
Provider.displayName = 'Provider';

export const App: React.FC = React.memo(() => {
  return (
    <Provider>
      {/* 基本CSS */}
      <CssBaseline />

      <NavBar />

      <MainRoute />
    </Provider>
  );
});
App.displayName = 'App';
