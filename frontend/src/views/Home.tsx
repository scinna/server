import {useToken} from "../utils/TokenProvider";
import React      from "react";
import { Redirect } from "react-router-dom";

export function Home() {
    const { userInfos } = useToken();

    if (userInfos) {
        return <Redirect to={{
            pathname: ""
        }} />
    }

    return <Redirect to={{
        pathname: '/login'
    }} />;
}