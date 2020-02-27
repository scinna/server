import React, {useReducer} from 'react';
import { Switch, Route } from 'react-router-dom';

import { AppContext, InitialState } from './context';
import dbReducer from './reducers/database';

import CardDatabase from './components/CardDatabase';
import CardScinna   from './components/CardScinna';
import CardFinale   from './components/CardFinale';
import CardIntro    from './components/CardIntro';
import CardEmail    from './components/CardEmail';
import CardUser     from './components/CardUser';

import LogoScinna from './assets/logo.png';


export default function() {
  return (
    // @ts-ignore
    <AppContext.Provider value={useReducer(dbReducer, InitialState)}>
      <div id="mainApp">
        <nav>
          <img src={LogoScinna} alt="Logo" />
        </nav>
        
        <div id="SiteContent">
          <Switch>
            <Route path="/" exact>
              <CardIntro />
            </Route>
            <Route path="/database">
              <CardDatabase />
            </Route>
            <Route path="/smtp">
              <CardEmail />
            </Route>
            <Route path="/scinna">
              <CardScinna />
            </Route>
            <Route path="/user">
              <CardUser />
            </Route>
            <Route path="/finale">
              <CardFinale />
            </Route>
          </Switch>
        </div>
      </div>
    </AppContext.Provider>
  );
}