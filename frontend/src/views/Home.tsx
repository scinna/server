import {useToken} from "../context/TokenProvider";
import React      from "react";
import { Redirect } from "react-router-dom";

export function Home() {
    const { userInfos } = useToken();

    if (userInfos) {
        return <Redirect to={{
            pathname: "/browse/" + userInfos.Name
        }} />
    }

    return <Redirect to={{
        pathname: '/login'
    }} />;
}