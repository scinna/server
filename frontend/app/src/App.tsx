import React, {useReducer, useEffect} from 'react';
import { Typography } from '@material-ui/core';
import { Switch, Route } from 'react-router-dom';

import { AppContext, CtxInitialState } from './context';
import MainReducer from './reducers';
import Menu from './components/Menu';
import IndexPage from './pages/IndexPage';

import { APIConfig } from './api/Config';
import { APICheckToken } from './api/Login';
import { setAxiosToken } from './api/Axios';


function App() {
  const [state, dispatch] = useReducer(MainReducer, CtxInitialState);

  useEffect(() => {
    APIConfig(dispatch);

    // @TODO: Not that good, calls this route just after login, even though data was retreived
    if (state.User.Token.length > 0) {
      setAxiosToken(state.User.Token);
      APICheckToken(dispatch, state.User.Token);
    }
  }, [state.User.Token]);

  return (
    // @ts-ignore
    <AppContext.Provider value={[state, dispatch]}>
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
  