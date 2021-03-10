import React                from 'react';
import ReactDOM             from 'react-dom';
import App                  from './App';
import {init as initTranslations}         from './translations';

import './assets/scss/Index.scss';
import TokenProvider        from "./utils/TokenProvider";
import ServerConfigProvider from "./utils/ServerConfigProvider";

initTranslations();

ReactDOM.render(
    <React.StrictMode>
        <ServerConfigProvider>
            <TokenProvider>
                <App/>
            </TokenProvider>
        </ServerConfigProvider>
    </React.StrictMode>,
    document.getElementById('root')
);
