import './assets/scss/App.scss';
import {useToken}                               from "./utils/TokenProvider";
import React, {ReactNode, useEffect}            from "react";
import {BrowserRouter, Redirect, Route, Switch} from "react-router-dom";
import {MuiThemeProvider as ThemeProvider}      from "@material-ui/core";
import Navigation                               from "./components/Navigation";
import {Register}                               from "./views/Register";
import {Login}                                  from "./views/Login";
import {Logout}                                 from "./views/Logout";
import {Account}                                from "./views/Account";
import {Browser}                                from "./views/Browser";
import {ShowPicture}                            from "./views/ShowPicture";
import {Home}                                   from "./views/Home";
import {createMuiTheme}                         from "@material-ui/core";

const AuthenticatedRoute = (node: ReactNode) => {
    const {isAuthenticated} = useToken();
    if (!isAuthenticated) {
        return <Redirect to={{pathname: '/login'}}/>
    }

    return node;
}

function App() {
    const theme = createMuiTheme({
        palette: {
            type: 'dark',
            primary: {
                main: '#87E7E1',
            },
        },
    });
    const {init} = useToken();

    // Yeah probably not what I'm supposed to do but meh it works for now
    useEffect(() => {
        init();
    }, [init])

    return (
        <ThemeProvider theme={theme}>
            <BrowserRouter basename={process.env.PUBLIC_URL}>
                <Navigation/>

                <Switch>
                    <Route exact path="/register">
                        <Register/>
                    </Route>

                    <Route exact path="/login">
                        <Login/>
                    </Route>

                    <Route exact path="/logout">
                        {AuthenticatedRoute(<Logout/>)}
                    </Route>

                    <Route exact path="/account">
                        {AuthenticatedRoute(<Account/>)}
                    </Route>

                    {/* Meh but react router seems to work only like this*/}
                    <Route path="/browse/:username/:path+">
                        <Browser/>
                    </Route>

                    <Route path="/browse/:username">
                        <Browser/>
                    </Route>

                    <Route path="/:pictureId">
                        <ShowPicture/>
                    </Route>

                    <Route path="/">
                        <Home/>
                    </Route>
                </Switch>
            </BrowserRouter>
        </ThemeProvider>
    );
}

export default App;
