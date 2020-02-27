import React from 'react';
import {Link, Redirect} from 'react-router-dom';

import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';

import { useStateValue } from '../context';
import SetScinnaSettings from '../api/Scinna';
import {actionUpdateScinna} from '../actions/scinna';


const useStyles = makeStyles(theme => ({
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
    },
}));

const initialState = {
    ConfigValid: false,
};

export default function() {
    const classes = useStyles();
    
    const [state, setState] = React.useState(initialState);

    //@ts-ignore
    const [global, dispatch] = useStateValue();

    const handleInputChange = (field: string) => (e: any) => {
        dispatch(actionUpdateScinna({[field]: e.currentTarget.value}))
    }

    const handleChange = (field: string) => (event: any) => {
        handleInputChange(field)({ currentTarget: event.target})
    };

    const submit = (e: any) => {
        e.preventDefault();

        SetScinnaSettings(global.Scinna)
            .then((r: any) => {
                setState({
                    ...state,
                    ConfigValid: r.data.IsValid,
                })
            })
            .catch((e: any) => {
                console.log(e)
            })

        return false;
    };


    // @TODO: window.location.protocol + "//" + window.location.hostname as default value
    return <div className="card above">
        { state.ConfigValid ? <Redirect to="/user" /> : null}
        <h4>About this server</h4>
        <form onSubmit={submit}>
            <div className="content centered-form">
                <p>This is really important. Please <a href="https://github.com/scinna/server/wiki/First-launch#scinna-settings" rel="noopener noreferrer" target="_blank">follow the docs</a> to understand each options.</p>
                <FormControl className={classes.formControl} fullWidth>
                    <InputLabel id="registration">Server registration</InputLabel>
                    <Select labelId="registration" id="registration" value={global.Scinna.Registration} onChange={handleChange("Registration")}>
                        <MenuItem value={"private"}>Private</MenuItem>
                        <MenuItem value={"public"}>Public</MenuItem>
                    </Select>
                </FormControl>
                <TextField id="scinna_header" label="IP Header" onChange={handleInputChange("IPHeader")} value={ global.Scinna.IPHeader } fullWidth />
                <TextField id="scinna_rate_limit" label="Rate limiting (Per 5 minutes)" onChange={handleInputChange("RateLimiting")} value={ global.Scinna.RateLimiting } fullWidth />
                <TextField id="scinna_path" label="Picture path" onChange={handleInputChange("PicturePath")} value={ global.Scinna.PicturePath } fullWidth />
                <TextField id="scinna_url" label="Web URL" onChange={handleInputChange("WebURL")} value={ global.Scinna.WebURL } fullWidth />
            </div>
            <div className="footer">
                <Link className="btn" to="/smtp">Back</Link>
                <Link className="btn" to="/user">Next</Link>
            </div>
        </form>
    </div>;
}