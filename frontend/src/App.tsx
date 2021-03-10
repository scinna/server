import './assets/scss/App.scss';
import {useToken}                           from "./utils/TokenProvider";
import React, {useEffect}                   from "react";
import {BrowserRouter, Switch, Route} from "react-router-dom";
import Navigation                           from "./components/Navigation";

function App() {
    const {init, userInfos} = useToken();
    useEffect(() => {
        init();
    }, [init])

    return (
        <BrowserRouter basename={process.env.PUBLIC_URL}>
            <Navigation />

            <Switch>
                <Route path="/logout">
                    <div>Logging-out</div>
                </Route>

                <Route path="/account">
                    <div>My account</div>
                </Route>

                <Route path="/">
                    <div>Home</div>
                </Route>
            </Switch>
        </BrowserRouter>
    );
}

export default App;
