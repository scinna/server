import {useToken} from "../context/TokenProvider";
import {Redirect} from "react-router-dom";

export function Logout() {
    useToken().logout();
    return <Redirect to={"/login"}/>;
}