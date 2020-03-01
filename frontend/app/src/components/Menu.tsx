import React, {useEffect} from 'react';
import {Link} from 'react-router-dom';

import { IconButton, Tooltip } from '@material-ui/core';
import MenuIcon from '@material-ui/icons/Menu';
import { makeStyles } from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';

import HomeIcon from '@material-ui/icons/Home';
import ProfilIcon from '@material-ui/icons/AccountBox';
import LogoutIcon from '@material-ui/icons/ExitToApp';

import {useStateValue} from '../context';
import { actionMenuToggle } from '../actions/MainActions';

import LogoScinna from '../assets/logo.png';
import '../assets/main.scss';

const useStyles = makeStyles(theme => ({
    drawerImg: {
        height: '3em',
        margin: '1em',
    }
}))

const MenuItem = (dispatch: any, icon: JSX.Element, text: string, to: string) => {
    return <ListItem button key={text} component={Link} to={to} onClick={() => dispatch(actionMenuToggle(false))}>
                <ListItemIcon>{icon}</ListItemIcon>
                <ListItemText primary={text}/>
            </ListItem>;
};

export default function() {
    const classes = useStyles();

    //@ts-ignore
    const [global, dispatch] = useStateValue();

    /**
     * This lets us close the menu when the user make the window grow again
     */
    const onResize = (evt: any) => {
        if (global.Main.menuOpened) {
            if (window.innerWidth > 900) {
                dispatch(actionMenuToggle(false));
            }
        }
    };

    useEffect(() => {
        window.addEventListener("resize", onResize);
        
        return function cleanup() {
        window.removeEventListener("resize", onResize);
        };
    });

    const userLoggedIn = global.User.Username.length > 0 && global.User.Token.length > 0;

    let menuLeft;
    let menuRight;
    let menuMobile;

    if (userLoggedIn) {
        menuLeft = [
            <li key="content"><Link to="/">My content</Link></li>,
            <li key="profile">
                <Tooltip title="My account" aria-label="My account">
                    <Link to="/me">{global.User.Username}</Link>
                </Tooltip>
            </li>
        ];

        menuRight = [
            <li key="home">
                <Tooltip title="Logout" aria-label="Logout">
                    <IconButton><LogoutIcon/></IconButton>
                </Tooltip>
            </li>,
        ];

        menuMobile = [
            MenuItem(dispatch, <HomeIcon />, "My content", "/"),
            MenuItem(dispatch, <ProfilIcon />, global.User.Username, "/me"),
            MenuItem(dispatch, <LogoutIcon />, "Logout", "/logout"), // @TODO: Just create an action that logs out
        ];
    } else {
        menuLeft = [
            <li key="home"><Link to="/">Home</Link></li>,
        ];

        menuRight = [
            <li key="home"><Link to="/">Login</Link></li>,
        ]
    }

    return <nav id="mainMenu">
        <IconButton aria-label="Open menu" className="hamburger" onClick={ () => dispatch(actionMenuToggle(true)) }>
            <MenuIcon />
        </IconButton>
        <Link to="/"><img src={LogoScinna} alt="Logo" /></Link>
        <ul className="menu">
            {menuLeft}
        </ul>

        <ul>
            {menuRight}
        </ul>

        <Drawer open={global.Main.menuOpened} onClose={ () => dispatch(actionMenuToggle(false)) }>
            <Link to="/"><img src={LogoScinna} alt="Logo" className={ classes.drawerImg } /></Link>
            <List>
                {menuMobile}
            </List>
        </Drawer>
    </nav>;
}