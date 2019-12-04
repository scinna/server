import React from 'react';
import {connect} from 'react-redux';
import {withRouter} from 'react-router-dom';

class Homepage extends React.Component {

    render() {

        let component;

        // Check if logged in
        if (!false) {
            component = <article>
                <h1>Scinna</h1>
                <h2>Picture sharing server</h2>
                <h3><a href="https://github.com/oxodao/scinna">Find yours on Github</a></h3>
                </article>
            } else {

            }

            return component;
        }

    }

    export default withRouter(
    connect(
        state => ({

    }),
    dispatch => ({

    })
)(Homepage));