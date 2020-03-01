import React from 'react';
import App from './App';

import { HashRouter } from 'react-router-dom';
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';
import { ThemeOptions } from '@material-ui/core/styles/createMuiTheme';

import ReactDOM from 'react-dom';

/**
 * @TODO: https://material-ui.com/customization/typography/
 *  Self-host fonts to not require CDN
 * Remove the dependencies in index.html
 */

const theme = createMuiTheme({
    palette: {
        type: 'dark',
        primary: { main: '#87E7E1', contrastText: '#fff' },
        secondary: { main: '#02968B', contrastText: '#fff' },
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
