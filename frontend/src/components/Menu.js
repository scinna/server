import React from 'react';
import {Link, withRouter} from 'react-router-dom'
import { connect } from 'react-redux';

import logo from '../assets/images/logo.png';
import '../assets/css/menu.scss'

class Menu extends React.Component {

    render() {
        let menuEnd = <div class="menu-end">
            <Link class="button" to="/logout">Logout</Link>
        </div>;

        if (false) {
            menuEnd = <div class="menu-end">
                            <Link class="button" to="/register"><strong>Sign up</strong></Link>
                            <Link class="button" to="/login">Log in</Link>
                        </div>;
        }

        return <nav>
                    <Link class="navbar-item" to="/">
                        <img className="brand" src={logo} alt="Main logo"/>
                    </Link>
                
                    <div id="menu-links">
                        <div class="menu-start">
                            <Link className="navbar-item" to="/">Home</Link>
                            <Link className="navbar-item" to="/">Documentation</Link>
                        </div>

                        {menuEnd}
                    </div>
                </nav>;

    }

}

export default withRouter(connect(
    state => ({

    }), 
    dispatch => ({

    })
)(Menu));