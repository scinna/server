import React from 'react';
import ReactDOM from 'react-dom';
import { HashRouter } from 'react-router-dom';
import App from './App';


import './assets/main.css';
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';
import { ThemeOptions } from '@material-ui/core/styles/createMuiTheme';


const theme = createMuiTheme({
    palette: {
        primary: { main: '#87E7E1' },
        secondary: { main: '#02968B' },
    }
  } as ThemeOptions);


ReactDOM.render(<HashRouter>
        <ThemeProvider theme={theme}><App /></ThemeProvider>
</HashRouter>, document.getElementById('root'));