import React from 'react';
import ReactDOM from 'react-dom';
import { HashRouter } from 'react-router-dom';
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';
import { ThemeOptions } from '@material-ui/core/styles/createMuiTheme';

import App from './App';


import './assets/main.css';


const theme = createMuiTheme({
    palette: {
        primary: { main: '#87E7E1' },
        secondary: { main: '#02968B' },
        text: {
          primary: "#fff",
          secondary: "rgba(255, 255, 255, .5)"
        },
        background: {
          paper: "var(--below-bg-color)"
        }
    }
  } as ThemeOptions);


ReactDOM.render(<HashRouter>
        <ThemeProvider theme={theme}><App /></ThemeProvider>
</HashRouter>, document.getElementById('root'));