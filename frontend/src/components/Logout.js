import React from 'react';
import {connect} from 'react-redux';
import {Link} from 'react-router-dom';

class Logout extends React.Component {

    componentDidMount() {
        // @TODO:send disconnect action
    }   

    render() {
        return <article>
            <h2>You've been logged out.</h2>
            <h3><Link to="/">Go back home</Link></h3>
        </article>;
    }

}

export default connect(state => {}, dispatch => {

})(Logout);