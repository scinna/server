import React      from 'react';
import {useToken} from "../utils/TokenProvider";

export function Profile() {
    const { userInfos } = useToken();

    return <div>Editing profile {userInfos?.Name}</div>
}