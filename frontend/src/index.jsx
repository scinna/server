import React         from 'react';
import ReactDOM      from 'react-dom';
import App           from './App';

import './assets/scss/Index.scss';
import TokenProvider from "./utils/TokenProvider";

ReactDOM.render(
  <React.StrictMode>
      <TokenProvider>
          <App />
      </TokenProvider>
  </React.StrictMode>,
  document.getElementById('root')
);
