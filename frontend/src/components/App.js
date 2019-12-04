import React     from 'react';
import Menu      from './Menu';
import Login     from './Login';
import Logout    from './Logout';
import SignUp    from './SignUp';
import Homepage  from './Homepage'
import {connect} from 'react-redux';

import { HashRouter as Router, Switch, Route } from 'react-router-dom'

import '../assets/css/index.scss';
import { bindActionCreators } from 'redux';
import { getTokenAction } from '../actions/AuthActions';


class App extends React.Component {

  componentDidMount() {
    // @TODO check if token already loaded
    this.props.loadToken();
  }  

  render() {
    return (<Router>
      <Menu/>
      <Switch>
        <Route exact path="/">
          <Homepage/>
        </Route>
        <Route path="/login">
          <Login />
        </Route>
        <Route path="/logout">
          <Logout />
        </Route>
        <Route path="/signup">
          <SignUp />
        </Route>
      </Switch>
    </Router>);
  }
}

export default connect(
  state => ({}),
  dispatch => ({
    loadToken: bindActionCreators(getTokenAction, dispatch)
  })
)(App);
