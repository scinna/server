import {useToken} from "../utils/TokenProvider";
import {Redirect} from "react-router-dom";

export function Logout() {
    useToken().logout();
    return <Redirect to={"/login"}/>;
}