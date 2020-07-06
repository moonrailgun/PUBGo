import React from 'react';
import ReactDom from 'react-dom';
import './utils/overrideMemo';
import { App } from './App';

ReactDom.render(<App />, document.querySelector('#app'));
