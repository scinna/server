import React from 'react';
import {Link, Redirect} from 'react-router-dom';

import TestMail from '../api/Mail';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Checkbox from '@material-ui/core/Checkbox';
import FormControlLabel from '@material-ui/core/FormControlLabel';

const useStyles = makeStyles(theme => ({
    fieldColor:{
        color: '#e4e4e5',
    }
}));

const initialState = {
    ConfigValid: false,
    MailDisabled: false,
    ButtonsEnabled: true,
};

export default function() {
    const classes = useStyles();
    /** 
     * @TODO: use Context to store the field inputs
     * so that changing page and going back keeps the values
     */
    
    const [state, setState] = React.useState(initialState);

    const [input, setInput] = React.useState({ MailDisabled: state.MailDisabled })
    const handleInputChange = (e: any) => setInput({
        ...input,
        [e.currentTarget.name]: e.currentTarget.value
    })

    const handleToggle = (event: any) => {
        setState({
            ...state,
            MailDisabled: event.target.checked,
        });
        setInput({
            ...input,
            MailDisabled: event.target.checked,
        })
    };

    const submit = (e: any) => {
        e.preventDefault();

        TestMail(input)
            .then((r: any) => {
                setState({
                    ...state,
                    ConfigValid: r.data.IsValid
                })
            })
            .catch((e: any) => {
                console.log(e)
            })

        return false;
    };

    let fields;
    if (state.MailDisabled) {
        fields = <div>
            <p>You chose to disable the emails.</p>
            <p>You will have to validate each account manually.</p>
            <p>You will also have to reset their forgotten passwords.</p>
        </div>;
    } else {
        fields = <div>
            <TextField id="smtp_host" name="smtp_host" label="Hostname" fullWidth onChange={handleInputChange} InputProps={{ className: classes.fieldColor }} />
            <TextField id="smtp_port" name="smtp_port" label="Port" fullWidth onChange={handleInputChange} InputProps={{ className: classes.fieldColor }} />
            <TextField id="smtp_username" name="smtp_username" label="Username" fullWidth onChange={handleInputChange} InputProps={{ className: classes.fieldColor }} />
            <TextField id="smtp_password" name="smtp_password" label="Password" fullWidth onChange={handleInputChange} InputProps={{ className: classes.fieldColor }} />
        </div>;
    }

    return <div className="card above">
        { state.ConfigValid ? <Redirect to="/user" /> : null}
        <h4>Email settings</h4>
        <form onSubmit={submit}>
            <div className="content centered-form">
                <p>Please fill the SMTP settings.</p>
                <FormControlLabel control={ <Checkbox checked={state.MailDisabled} onChange={handleToggle} value="MailDisabled" />} label="Disable emails ?" />
                {fields}
            </div>
            <div className="footer">
                <Link className="btn" to="/database">Back</Link>
                <input type="submit" className="btn" value="Next" />
            </div>
        </form>
    </div>;
}