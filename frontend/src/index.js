import * as serviceWorker from './serviceWorker';
import { Provider }       from 'react-redux';
import ReactDOM           from 'react-dom';
import reducers           from './reducers';
import React              from 'react';
import App                from './components/App.js';
import vsaga              from './saga';

import localStorageMiddleware                   from './middleware/LocalStorageMiddleware';
import createSagaMiddleware                     from 'redux-saga';
import {createStore, applyMiddleware, compose}  from 'redux';

const sagaMiddleware = createSagaMiddleware();

let middleware       = [sagaMiddleware, localStorageMiddleware];

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const store            = createStore(reducers, composeEnhancers(applyMiddleware(...middleware)));

sagaMiddleware.run(vsaga);

ReactDOM.render(<Provider store={store}>
                    <App />
                </Provider>, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
