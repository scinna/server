import {Link}     from "react-router-dom";
import React      from "react";
import {useToken} from "../utils/TokenProvider";

import "../assets/scss/Navigation.scss";

export default function Navigation() {
    const {userInfos} = useToken();

    return (
        <nav>
            <img alt="Logo" src="/api/server/logo"/>
            <ul>
                <li><Link to="/">Home</Link></li>

                {
                    (!userInfos) && <>
                        <li><Link to="/login">Login</Link></li>
                        <li><Link to="/register">Register</Link></li>
                    </>
                }

                {
                    (userInfos) && <>
                        <li><Link to="/account">{userInfos.Name}</Link></li>
                        <li><Link to="/logout">Logout</Link></li>
                    </>
                }
            </ul>
        </nav>
    );
}