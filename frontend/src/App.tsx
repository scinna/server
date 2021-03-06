import './assets/scss/App.scss';
import {useToken}                               from "./context/TokenProvider";
import React, {ReactNode, useEffect}            from "react";
import {BrowserRouter, Redirect, Route, Switch} from "react-router-dom";
import {MuiThemeProvider as ThemeProvider}      from "@material-ui/core";
import Navigation                               from "./components/Navigation";
import BrowserProvider                          from "./context/BrowserProvider";
import {Register}                               from "./views/Register";
import {Login}                                  from "./views/Login";
import {Logout}                                 from "./views/Logout";
import {Account}                                from "./views/Account";
import {Browser}                                from "./views/Browser";
import {ShowMedia}                              from "./views/ShowMedia";
import {Home}                                   from "./views/Home";
import {createMuiTheme}                         from "@material-ui/core";
import {ServerSettings}                         from "./views/ServerSettings";
import {ValidateAccount}                        from "./views/Validate";
import {LinkShortnener}                         from "./views/LinkShortener";
import {PasswordReset}                          from "./views/PasswordReset";
import ShortenLinkProvider                      from "./context/ShortenLinkProvider";
import ModalProvider                            from "./context/ModalProvider";

const AuthenticatedRoute = (node: ReactNode) => {
    const {isAuthenticated} = useToken();
    if (!isAuthenticated) {
        return <Redirect to={{pathname: '/login'}}/>
    }

    return node;
}

/** Role will be used when a real role system will be implemented **/
const AdminAuthenticatedRoute = (node: ReactNode, role: string) => {
    const {isAuthenticated, userInfos} = useToken();
    if (!isAuthenticated) {
        return <Redirect to={{pathname: '/login'}}/>
    }

    if (!userInfos?.IsAdmin) {
        return <Redirect to={{pathname: '/account'}}/>
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
            background: {
                default: '#222d32',
                paper: '#1e282c',
            }
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
                <ModalProvider>
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

                        <Route path="/validate/:valCode">
                            <ValidateAccount/>
                        </Route>

                        <Route path="/forgotten_password/:valCode">
                            <PasswordReset/>
                        </Route>

                        <Route exact path="/account">
                            {AuthenticatedRoute(<Account/>)}
                        </Route>

                        <Route exact path="/shortener">
                            <ShortenLinkProvider>
                                {AuthenticatedRoute(<LinkShortnener/>)}
                            </ShortenLinkProvider>
                        </Route>

                        <Route exact path="/admin">
                            {AdminAuthenticatedRoute(<ServerSettings/>, 'ROLE_ADMIN')}
                        </Route>

                            {/* Meh but react router seems to work only like this*/}
                            <Route path="/browse/:username/:path+">
                                <BrowserProvider>
                                    <Browser/>
                                </BrowserProvider>
                            </Route>

                            <Route path="/browse/:username">
                                <Browser/>
                            </Route>


                        <Route path="/:pictureId">
                            <ShowMedia/>
                        </Route>

                        <Route path="/">
                            <Home/>
                        </Route>
                    </Switch>
                </ModalProvider>
            </BrowserRouter>
        </ThemeProvider>
    );
}

export default App;
