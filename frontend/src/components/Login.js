import React from 'react'
import {withRouter} from 'react-router-dom';

import '../assets/css/login.scss';

class Login extends React.Component {

    render() {
        return <article>
                <div class="card">
                    <header>Login</header>

                    <div class="card-content" style={{ width:"350px"}}>
                            <input class="input" type="text" placeholder="Username" />
                            <input class="input" type="password" placeholder="Password"/>
                    </div>

                    <footer class="card-footer container">
                        <button class="card-footer-item button">Login</button>
                    </footer>
                </div>
        </article>
    }

}

export default withRouter(Login);