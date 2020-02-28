import React from 'react';
import {Link, Redirect} from 'react-router-dom';

import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Snackbar from '@material-ui/core/Snackbar';
import IconButton from '@material-ui/core/IconButton';
import MuiAlert from '@material-ui/lab/Alert';
import CloseIcon from '@material-ui/icons/Close';

import TestDB from '../api/Database';
import { useStateValue } from '../context';
import { actionUpdateDatabase } from '../actions/database';

const useStyles = makeStyles(theme => ({
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
    },
}));

const initialState = {
    ConfigValid: false,
    ButtonsEnabled: true,
    SnackbarOpened: false,
    SnackbarMessage: "",
};

/**
 * @TODO: Lock the buttons on all interfaces until the user receives the response from the HTTP request
 */
export default function() {
    const classes = useStyles();
    const [state, setState] = React.useState(initialState);

    //@ts-ignore
    const [global, dispatch] = useStateValue();

    const handleCloseSnack = (e: any) => {
        setState({
            ...state,
            SnackbarMessage: "",
            SnackbarOpened: false,
        })
    };

    const handleInputChange = (field: string) => (e: any) => {
        if (e.currentTarget.type === "number") {
            let val = parseInt(e.currentTarget.value)
            dispatch(actionUpdateDatabase({ [field]: val }))   
        } else {
            dispatch(actionUpdateDatabase({ [field]: e.currentTarget.value }))   
        }
    }

    const submit = (e: any) => {
        e.preventDefault();

        // @TODO Find a way to use useEffect (?)
        // Not working because not in React component but in submit function
        TestDB(global.Database)
            .then((r: any) => {
                setState({
                    ...state,
                    ConfigValid: r.data.IsValid,
                    SnackbarOpened: !r.data.IsValid || r.status !== 200,
                    SnackbarMessage: r.data.Message,    
                })
            })
            .catch((e: any) => {
                setState({
                    ...state,
                    ConfigValid: e.response.data.IsValid,
                    SnackbarMessage: e.response.data.Message,
                    SnackbarOpened: true,
                })
            })

        return false;
    };

    let fields;
    if (global.Database.Dbms === 'sqlite3') {
        fields = <div>
            <TextField id="db_path" name="db_path" label="File path" required fullWidth onChange={handleInputChange("Path")} value={global.Database.Path} />
        </div>;
    } else if (global.Database.Dbms === 'pgsql' || global.Database.Dbms === 'mysql') {
        fields = <div>
            <div style={{display: 'flex'}}>
                <TextField id="db_host" name="db_host" label="Hostname" required onChange={handleInputChange("Hostname")} value={ global.Database.Hostname } />
                <TextField id="db_port" name="db_port" label="Port" required type="number" inputProps={{ min: "0" }} onChange={handleInputChange("Port")} value={ global.Database.Port } />
            </div>
            <TextField id="db_username" name="db_username" required label="Username" fullWidth onChange={handleInputChange("Username")} value={global.Database.Username} />
            <TextField id="db_password" name="db_password" required label="Password" fullWidth onChange={handleInputChange("Password")} value={global.Database.Password} type="password" />
            <TextField id="db_database" name="db_database" required label="Database" fullWidth onChange={handleInputChange("Database")} value={global.Database.Database} />
        </div>;
    } else {
        fields = <div>
            <p>This shouldn't happen as no other DBMS are supported for now.</p>
        </div>;
    }

    return <div className="card above">
        { state.ConfigValid ? <Redirect to="/smtp" /> : null}
        <h4>Database settings</h4>
        <form onSubmit={submit}>
            <div className="content centered-form">
                <p>Please choose your database software.</p>
                <FormControl className={classes.formControl} fullWidth>
                    <InputLabel id="SelectDBMS">Choose your DBMS</InputLabel>
                    <Select labelId="SelectDBMS" id="dbms" name="dbms" value={global.Database.Dbms} onChange={handleInputChange("Dbms")}>
                        <MenuItem value={"pgsql"}>Postgres</MenuItem>
                        <MenuItem value={"mysql"}>Mysql</MenuItem>
                        <MenuItem value={"sqlite3"}>Sqlite3</MenuItem>
                    </Select>
                </FormControl>

                {fields}
            </div>
            <div className="footer">
                <Link className="btn" to="/">Back</Link>
                <input type="submit" className="btn" value="Next" />
            </div>
        </form>

        <Snackbar 
            anchorOrigin={{ vertical: 'bottom', horizontal: 'center', }}
            open={state.SnackbarOpened}
            autoHideDuration={6000}
            onClose={handleCloseSnack}
            action={
                <React.Fragment>
                    <IconButton size="small" aria-label="close" color="inherit" onClick={handleCloseSnack}>
                        <CloseIcon fontSize="small" />
                    </IconButton>
                </React.Fragment>
            }
      >
            <MuiAlert elevation={6} variant="filled" severity="error">{state.SnackbarMessage}</MuiAlert>
        </Snackbar> 
    </div>;
}