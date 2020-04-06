import React, {useReducer, useEffect} from 'react';
import { Typography } from '@material-ui/core';
import { Switch, Route } from 'react-router-dom';

import { AppContext, CtxInitialState } from './context';
import MainReducer from './reducers';
import Menu from './components/Menu';
import IndexPage from './pages/IndexPage';
import ProfilePage from './pages/ProfilePage';

import { AxiosMiddlishware } from './api/Axios';
import MediaViewer from './components/MediaViewer';


function App() {
  const [state, dispatch] = useReducer(MainReducer, CtxInitialState);

  AxiosMiddlishware(dispatch);

  return (
    // @ts-ignore
    <AppContext.Provider value={[state, dispatch]}>
      <div id="mainApp">
        <Menu/>
        <Switch>
          <Route path="/media/:pictID">
            <MediaViewer/>
          </Route>
          <Route path="/me">
            <ProfilePage/>
          </Route>

          <Route exact={false} path="/">
            <IndexPage/>
          </Route>
        </Switch>
      </div>
    </AppContext.Provider>);
  }
  
  export default App;
  