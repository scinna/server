import React from 'react';
import {Link, Redirect} from 'react-router-dom';

import TestMail from '../api/Mail';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Checkbox from '@material-ui/core/Checkbox';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Snackbar from '@material-ui/core/Snackbar';
import IconButton from '@material-ui/core/IconButton';
import MuiAlert from '@material-ui/lab/Alert';
import CloseIcon from '@material-ui/icons/Close';
import ValidIcon from '@material-ui/icons/Check';

import { useStateValue } from '../context';
import {actionUpdateSmtp} from '../actions/smtp';

const useStyles = makeStyles(theme => ({
    fieldColor:{
        color: '#e4e4e5',
    }
}));

const initialState = {
    ConfigValid: false,
    ButtonsEnabled: true,
    SnackbarOpened: false,
    SnackbarMessage: '',
    SnackbarError: false,
};

export default function() {
    const classes = useStyles();    
    const [state, setState] = React.useState(initialState);

    //@ts-ignore
    const [global, dispatch] = useStateValue();

    const handleInputChange = (field: string) => (e: any) => {
        dispatch(actionUpdateSmtp({[field]: e.currentTarget.value}))
    }

    const handleToggle = (field: string) => (event: any) => {
        dispatch(actionUpdateSmtp({[field]: !event.target.checked}))
    };

    
    const handleCloseSnack = (e: any) => {
        if (!(e instanceof MouseEvent)) {
            setState({
                ...state,
                SnackbarOpened: false,
            })
        }
    };

    const handleReceived = (e: any) => {
        setState({
            ...state,
            ConfigValid: true
        })
    }

    const submit = (e: any) => {
        e.preventDefault();

        TestMail(global.Smtp)
            .then((r: any) => {
                let message = r.data.IsValid ? "Mail sent. Did you received it?" : "Something went wrong.";
                setState({
                    ...state,
                    //ConfigValid: r.data.IsValid,
                    SnackbarOpened: true,
                    SnackbarMessage: message,
                    SnackbarError: !r.data.IsValid
                })
            })
            .catch((e: any) => {
                console.log(e)
            })

        return false;
    };

    /**
     * @TODO: Add a button to send a test mail
     *        Add a popup (Or snackbar) to ask the user if he received the email
     */

    let fields;
    if (!global.Smtp.Enabled) {
        fields = <div>
            <p>You chose to disable the emails.</p>
            <p>You will have to validate each account manually.</p>
            <p>Users will also not be able to recover their account if they forget their passwords.</p>
        </div>;
    } else {
        fields = <div>
            <p>This only support STARTTLS for now.</p>
            <TextField id="smtp_host" name="smtp_host" label="Hostname" fullWidth onChange={handleInputChange("Hostname")} value={ global.Smtp.Hostname } InputProps={{ className: classes.fieldColor }} />
            <TextField id="smtp_port" name="smtp_port" label="Port" fullWidth onChange={handleInputChange("Port")} value={ global.Smtp.Port } InputProps={{ className: classes.fieldColor }} />
            <TextField id="smtp_username" name="smtp_username" label="Username" fullWidth onChange={handleInputChange("Username")} value={ global.Smtp.Username } InputProps={{ className: classes.fieldColor }} />
            <TextField id="smtp_password" name="smtp_password" type="password" label="Password" fullWidth onChange={handleInputChange("Password")} value={ global.Smtp.Password } InputProps={{ className: classes.fieldColor }} />
            <TextField id="smtp_sender" name="smtp_sender" label="Sender" fullWidth onChange={handleInputChange("Sender")} value={ global.Smtp.Sender } InputProps={{ className: classes.fieldColor }} />
            <TextField id="smtp_receiver" name="smtp_receiver" label="Test receiver" fullWidth onChange={handleInputChange("TestReceiver")} value={ global.Smtp.TestReceiver } InputProps={{ className: classes.fieldColor }} />
        </div>;
    }

    let btText = global.Smtp.Enabled ? "Send test mail" : "Next";

    return <div className="card above">
        { state.ConfigValid ? <Redirect to="/scinna" /> : null}
        <h4>Email settings</h4>
        <form onSubmit={submit}>
            <div className="content centered-form">
                <p>Please fill the SMTP settings.</p>
                <FormControlLabel control={ <Checkbox checked={!global.Smtp.Enabled} onChange={handleToggle("Enabled")} value="MailDisabled" />} label="Disable emails ?" />
                {fields}
            </div>
            <div className="footer">
                <Link className="btn" to="/database">Back</Link>
                <input type="submit" className="btn" value={btText} />
            </div>
        </form>
        
        <Snackbar 
            anchorOrigin={{ vertical: 'bottom', horizontal: 'center', }}
            open={state.SnackbarOpened}
            onClose={handleCloseSnack}
            >
            <MuiAlert elevation={6} 
                variant="filled"
                severity={state.SnackbarError ? "error" : "info"} 
                action={
                    <React.Fragment>
                        <IconButton size="small" aria-label="close" color="inherit" onClick={handleReceived}>
                            <ValidIcon fontSize="small" />
                        </IconButton>
                        <IconButton size="small" aria-label="close" color="inherit" onClick={handleCloseSnack}>
                            <CloseIcon fontSize="small" />
                        </IconButton>
                    </React.Fragment>
                }>
                    {state.SnackbarMessage}
            </MuiAlert>
        </Snackbar>
    </div>;
}