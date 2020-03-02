import React from 'react';

import {Grid, Typography, makeStyles } from '@material-ui/core';
import TextField from '@material-ui/core/TextField';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Box from '@material-ui/core/Box';
import Button from '@material-ui/core/Button';

import {useStateValue} from '../context';

import '../assets/Login.scss';

const useStyles = makeStyles(theme => ({
    container: {
        display: 'block'
    },
    text: {
        marginTop: '2em',
        display: 'block',
    },
    submitForm: {
        marginTop: '4em',
    },
    loginForm: {
        display: 'flex',
        flexDirection: 'column',
        margin: 'auto',
    }
}))

const TabPanel = (props: any) => {
    const { children, value, index, ...others } = props;

    return <Box p={3} {...others}>
        {value === index && {...children}}
    </Box>
}

export default function() {
    const classes = useStyles();
    const [tab, setTab] = React.useState(1);

    //@ts-ignore
    const [global, dispatch] = useStateValue();
    
    const handleChange = (event: any, tab: number) => {
        setTab(tab);
    };

    return <Grid container>
        <Grid container item xs={12} sm={6} className={'LoginPadding ' + classes.container}>
            <Typography color="textPrimary" variant="h3" component="h1">
                Scinna
            </Typography>
            <Typography color="textSecondary" component="p" className={classes.text}>
                This is a Scinna instance.
            </Typography>
            <Typography color="textSecondary" component="p" className={classes.text}>
                Scinna is an open-source, self-hosted screenshot and picture server. Find more about it at <a className="pretty-link" href="https://scinna.app/">https://scinna.app</a>
            </Typography>
        </Grid>
        
        <Grid container item xs={12} sm={6} className={'LoginPadding ' + classes.container}>
            <Tabs value={tab} indicatorColor="primary" textColor="primary" centered onChange={handleChange} aria-label="Login and registration panels">
                <Tab label="Registration"/>
                <Tab label="Login"/>
            </Tabs>

            <TabPanel value={tab} index={0}>
                <form className={`FormMaxWidth ${classes.loginForm}`}>
                    <TextField id="username" label="Username" required/>
                    <TextField id="username" label="Email" type="email" required/>
                    <TextField id="password" label="Password" type="password" required/>
                    <TextField id="password2" label="Repeat password" type="password" required/>
                    { global.Config.Registration === "private" && <TextField id="InviteCode" label="Invitation code" type="text" required />}
 
                    <Button variant="contained" color="secondary" className={classes.submitForm} type="submit">Register</Button>
                </form>
            </TabPanel>

            <TabPanel value={tab} index={1}>
                <form className={`FormMaxWidth ${classes.loginForm}`}>
                    <TextField id="username" label="Username" required/>
                    <TextField id="password" label="Password" required/>
                    { global.Config.EmailAvailable && <a>Forgotten password?</a>}
 
                    <Button variant="contained" color="secondary" className={classes.submitForm} type="submit"> Login </Button>
                </form>
            </TabPanel>
        </Grid>
    </Grid>;
}