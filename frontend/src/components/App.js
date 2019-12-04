import React    from 'react';
import Menu     from './Menu';
import Login    from './Login';
import SignUp   from './SignUp';
import Homepage from './Homepage'

import { HashRouter as Router, Switch, Route } from 'react-router-dom'

import '../assets/css/index.scss';

function App() {
  return (<Router>
    <Menu/>
    <Switch>
      <Route exact path="/">
        <Homepage/>
      </Route>
      <Route path="/login">
        <Login />
      </Route>
      <Route path="/signup">
        <SignUp />
      </Route>
    </Switch>
  </Router>);
}

export default App;
