import React from 'react';
import {Link, Redirect} from 'react-router-dom';

import MuiAlert from '@material-ui/lab/Alert';
import Snackbar from '@material-ui/core/Snackbar';
import TextField from '@material-ui/core/TextField';

import { useStateValue } from '../context';
import {actionUpdateUser} from '../actions/user';
import CreateUserAccount from '../api/User';

const initialState = {
    ConfigValid: false,
    PwdRepeat: "",
    SnackbarOpened: false,
};

export default function() {
    const [state, setState] = React.useState(initialState);

    //@ts-ignore
    const [global, dispatch] = useStateValue();

    const handleCloseSnack = (e: any) => {
        setState({
            ...state,
            SnackbarOpened: false,
        })
    };

    const handleInputChange = (field: string) => (e: any) => {
        dispatch(actionUpdateUser({[field]: e.currentTarget.value}))
    }

    const handlePwdRepeatChange = (e: any) => {
        setState({
            ...state,
            PwdRepeat: e.currentTarget.value,
        })
    }

    const callAPI = () => {
        CreateUserAccount(global.User)
            .then((r: any) => {
                setState({
                    ...state,
                    ConfigValid: true,
                })
            })
            .catch((e: any) => {
                console.log(e)
            })
    }

    const submit = (e: any) => {
        e.preventDefault();

        if (state.PwdRepeat === global.User.Password) {
            callAPI();
            return false;
        }

        setState({
            ...state,
            SnackbarOpened: true,
        })

        return false;
    };

    return <div className="card above">
        { state.ConfigValid ? <Redirect to="/finale" /> : null}
        <h4>Create your account</h4>
        <form onSubmit={submit}>
            <div className="content">
                <p>Creating the admin account.</p>
                <TextField name="user_name" label="Username" onChange={handleInputChange("Username")} value={ global.User.Username } required fullWidth />
                <TextField name="user_mail" label="Email" onChange={handleInputChange("Email")} value={ global.User.Email } required fullWidth />
                <TextField name="user_pass" label="Password" type="password" onChange={handleInputChange("Password")} value={ global.User.Password } required fullWidth />
                <TextField name="user_pwd2" label="Repeat password" type="password" onChange={handlePwdRepeatChange} value={ state.PwdRepeat } required fullWidth />
            </div>
            <div className="footer">
                <Link className="btn" to="/scinna">Back</Link>
                <input type="submit" className="btn" value="Next" />
            </div>
        </form>

        <Snackbar 
            anchorOrigin={{ vertical: 'bottom', horizontal: 'center', }}
            open={state.SnackbarOpened}
            autoHideDuration={3000}
            onClose={handleCloseSnack}>
            <MuiAlert elevation={6} variant="filled" severity="error">The passwords do not matches !</MuiAlert>
        </Snackbar> 
    </div>;
}