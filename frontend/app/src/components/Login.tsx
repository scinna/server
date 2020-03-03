import React from 'react';

import {Grid, Typography, makeStyles } from '@material-ui/core';
import TextField from '@material-ui/core/TextField';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Box from '@material-ui/core/Box';
import Button from '@material-ui/core/Button';
import Snackbar from '@material-ui/core/Snackbar';
import IconButton from '@material-ui/core/IconButton';
import {Alert, AlertTitle} from '@material-ui/lab';
import CloseIcon from '@material-ui/icons/Close'; 

import {useStateValue} from '../context';
import { APILogin } from '../api/Login';

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

const initialState: any = {
    Snackbar: {
        Opened: false,
        Severity: 'error',
        Message: '',
    },
    Registration: {
        Username: "",
        Email: "",
        Password: "",
        Password2: "",
    },
    Login: {
        Username: "",
        Password: "",
    }
}

export default function() {
    const classes = useStyles();
    const [tab, setTab] = React.useState(1);
    const [forms, setForms] = React.useState(initialState);

    //@ts-ignore
    const [global, dispatch] = useStateValue();
    
    const handleChange = (event: any, tab: number) => {
        setTab(tab);
    };
    
    const handleCloseSnack = (e: any) => {
        setForms({
            ...forms,
            Snackbar: {
                Opened: false,
                Severity: 'error',
                Message: '',
            }
        })
    };

    const handleInputChangeRegister = (e: any) => {
        const elt = e.currentTarget;
        setForms({ ...forms, Registration: { ...forms.Registration, [elt.id]: elt.value }})
    }

    const handleInputChangeLogin = (e: any) => {
        const elt = e.currentTarget;
        setForms({ ...forms, Login: { ...forms.Login, [elt.id]: elt.value }})
    }

    const registerSubmit = (event: any) => {
        if (forms.Registration.Password !== forms.Registration.Password2) {
            setForms({
                ...forms,
                Snackbar: {
                    Opened: true,
                    Severity: 'error',
                    Message: 'Passwords doesn\'t match!',
                }
            })
        } else {
            // @TODO: API Request
        }
    }

    const loginSubmit = (event: any) => {
        APILogin(dispatch, forms.Login, (resp: any) => {
            setForms({
                ...forms,
                Snackbar: {
                    Opened: true,
                    ...resp
                }
            })
        })
    }

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

            { 
            /**
             * The thing with InputLabelProps disable the small asterisk without removing the required prop
             **/ 
            }
            <TabPanel value={tab} index={0}>
                <form className={`FormMaxWidth ${classes.loginForm}`} onSubmit={registerSubmit}>
                    <TextField id="Username" label="Username" onChange={handleInputChangeRegister} value={forms.Registration.Username} required InputLabelProps={{ required: false }}/>
                    <TextField id="Email" label="Email" type="email" onChange={handleInputChangeRegister} value={forms.Registration.Email} required InputLabelProps={{ required: false }}/>
                    <TextField id="Password" label="Password" type="password" onChange={handleInputChangeRegister} value={forms.Registration.Password} required InputLabelProps={{ required: false }}/>
                    <TextField id="Password2" label="Repeat password" type="password" onChange={handleInputChangeRegister} value={forms.Registration.Password2} required InputLabelProps={{ required: false }}/>
                    { global.Config.Registration === "InviteCode" && <TextField id="invite" onChange={handleInputChangeRegister} value={forms.Registration.InviteCode} label="Invitation code" type="text" required InputLabelProps={{ required: false }} />}
 
                    <Button variant="contained" color="secondary" className={classes.submitForm} type="submit">Register</Button>
                </form>
            </TabPanel>

            <TabPanel value={tab} index={1}>
                <form className={`FormMaxWidth ${classes.loginForm}`} onSubmit={loginSubmit}>
                    <TextField id="Username" label="Username" onChange={handleInputChangeLogin} value={forms.Login.Username} required InputLabelProps={{ required: false }}/>
                    <TextField id="Password" label="Password" onChange={handleInputChangeLogin} value={forms.Login.Password} required InputLabelProps={{ required: false }} type="password" />
                    { global.Config.EmailAvailable && <a>Forgotten password?</a>}
 
                    <Button variant="contained" color="secondary" className={classes.submitForm} type="submit">Login</Button>
                </form>
            </TabPanel>
        </Grid>

        <Snackbar anchorOrigin={{ vertical: 'bottom', horizontal: 'center', }} open={forms.Snackbar.Opened}
            autoHideDuration={6000} onClose={handleCloseSnack}
            action={
                <React.Fragment>
                    <IconButton size="small" aria-label="close" color="inherit" onClick={handleCloseSnack}>
                        <CloseIcon fontSize="small" />
                    </IconButton>
                </React.Fragment>
            }
      >
            {/*
            // @ts-ignore */}
            <Alert elevation={6} 
                severity={forms.Snackbar.Severity} 
                action={
                    <React.Fragment>
                        <IconButton size="small" aria-label="close" color="inherit" onClick={handleCloseSnack}>
                            <CloseIcon fontSize="small" />
                        </IconButton>
                    </React.Fragment>
                }>
                    {forms.Snackbar.Message}
            </Alert>
        </Snackbar> 
    </Grid>;
}