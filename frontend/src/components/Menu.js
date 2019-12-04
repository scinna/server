import React from 'react';
import {Link, withRouter} from 'react-router-dom';

import logo from '../assets/images/logo.png';
import '../assets/css/menu.scss'

class Menu extends React.Component {

    render() {
        return <nav>
                    <Link class="navbar-item" to="/">
                        <img class="brand" src={logo} alt="Main logo"/>
                    </Link>
                
                    <div id="menu-links">
                        <div class="menu-start">
                            <Link className="navbar-item" to="/">Home</Link>
                            <Link className="navbar-item" to="/">Documentation</Link>
                        </div>

                        <div class="menu-end">
                            <Link class="button" to="/register"><strong>Sign up</strong></Link>
                            <Link class="button" to="/login">Log in</Link>
                        </div>
                    </div>
                </nav>;

    }

}

export default withRouter(Menu);