import React, {useReducer} from 'react';
import { Typography } from '@material-ui/core';
import { Switch, Route } from 'react-router-dom';

import { AppContext, InitialState } from './context';
import MainReducer from './reducers';
import Menu from './components/Menu';
import IndexPage from './pages/IndexPage';


function App() {
  return (
    // @ts-ignore
    <AppContext.Provider value={useReducer(MainReducer, InitialState)}>
      <div id="mainApp">
        <Menu/>
        <Switch>
          <Route path="/" exact>
            <IndexPage/>
          </Route>
          <Route path="/pictures/:pictID">
            <Typography>Seeing a picture</Typography>
          </Route>
          <Route path="/me">
            <Typography>Seeing my profile</Typography>
          </Route>
        </Switch>
      </div>
    </AppContext.Provider>);
  }
  
  export default App;
  