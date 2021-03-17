import {Link}                                                       from "react-router-dom";
import React, {useState}                                            from "react";
import {useToken}                                                   from "../utils/TokenProvider";
import {Button, Drawer, List, ListItem, ListItemIcon, ListItemText} from "@material-ui/core";
import i18n                                                         from 'i18n-js';
import {
    Home as HomeIcon,
    Person as PersonIcon,
    HowToReg as RegisterIcon,
    ExitToApp as LogoutIcon,
    Settings as SettingsIcon,
    Menu as MenuIcon
}                                                                   from "@material-ui/icons";

import styles from "../assets/scss/Navigation.module.scss";

export default function Navigation() {
    const [isOpened, setOpened] = useState<boolean>(false);
    const {userInfos} = useToken();

    return (
        <nav>
            <img alt="Logo" src="/api/server/logo"/>
            <ul>
                <li><Link to="/">{i18n.t('menu.home')}</Link></li>

                {
                    (!userInfos) && <>
                        <li><Link to="/login">{i18n.t('menu.login')}</Link></li>
                        <li><Link to="/register">{i18n.t('menu.register')}</Link></li>
                    </>
                }

                {
                    (userInfos) && <>
                        <li><Link to="/account">{userInfos.Name}</Link></li>
                        {
                            (userInfos.IsAdmin)
                            &&
                            <li><Link to="/admin">{i18n.t('menu.server')}</Link></li>
                        }
                        <li><Link to="/logout">{i18n.t('menu.logout')}</Link></li>
                    </>
                }
            </ul>

            <div className={styles.Hamburger}>
                <Button onClick={() => setOpened(true)}>
                    <MenuIcon/>
                </Button>
            </div>

            <Drawer anchor="left" open={isOpened} onClose={() => setOpened(false)}>
                <List>
                    <ListItem component={Link} to="/" onClick={() => setOpened(false)} button>
                        <ListItemIcon><HomeIcon/></ListItemIcon>
                        <ListItemText primary={i18n.t('menu.home')}/>
                    </ListItem>
                    {
                        (!userInfos) && <>
                            <ListItem component={Link} to="/" onClick={() => setOpened(false)} button>
                                <ListItemIcon><PersonIcon/></ListItemIcon>
                                <ListItemText primary={i18n.t('menu.login')}/>
                            </ListItem>
                            <ListItem component={Link} to="/register" onClick={() => setOpened(false)} button>
                                <ListItemIcon><RegisterIcon/></ListItemIcon>
                                <ListItemText primary={i18n.t('menu.register')}/>
                            </ListItem>
                        </>
                    }
                    {
                        (userInfos) && <>
                            <ListItem component={Link} to="/account" onClick={() => setOpened(false)} button>
                                <ListItemIcon><PersonIcon/></ListItemIcon>
                                <ListItemText primary={userInfos.Name}/>
                            </ListItem>
                            {
                                userInfos.IsAdmin
                                &&
                                <ListItem component={Link} to="/admin" onClick={() => setOpened(false)} button>
                                    <ListItemIcon><SettingsIcon/></ListItemIcon>
                                    <ListItemText primary={i18n.t('menu.server')}/>
                                </ListItem>
                            }
                            <ListItem component={Link} to="/logout" onClick={() => setOpened(false)} button>
                                <ListItemIcon><LogoutIcon/></ListItemIcon>
                                <ListItemText primary={i18n.t('menu.logout')}/>
                            </ListItem>
                        </>
                    }
                </List>
            </Drawer>
        </nav>
    );
}