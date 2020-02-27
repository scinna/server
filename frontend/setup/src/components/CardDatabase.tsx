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

/**
 * When the user click on Next, the app should send the data to the /database endpoint to test them
 * If the database doesn't work or it does not have enough rights, the app mustn't let the user process
 * forward.
 */

const useStyles = makeStyles(theme => ({
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
    },
    fieldColor:{
        color: '#e4e4e5',
    }
}));

const initialState = {
    ConfigValid: false,
    ButtonsEnabled: true,
    SnackbarOpened: false,
};

export default function() {
    const classes = useStyles();
    const [state, setState] = React.useState(initialState);

    // @TODO: Remove typescript >:(

    //@ts-ignore
    const [global, dispatch] = useStateValue();

    const handleCloseSnack = (e: any) => {
        setState({
            ...state,
            SnackbarOpened: false,
        })
    };

    const handleInputChange = (field: string) => (e: any) => {
        dispatch(actionUpdateDatabase( { [field]: e.target.value }))
    }

    const submit = (e: any) => {
        e.preventDefault();

        TestDB(global.Database)
            .then((r: any) => {
                setState({
                    ...state,
                    ConfigValid: r.data.IsValid,
                    SnackbarOpened: !r.data.IsValid,
                })
            })
            .catch((e: any) => {
                console.log(e)
            })

        return false;
    };

    let fields;
    if (global.Database.Dbms === 'sqlite3') {
        fields = <div>
            <TextField id="db_path" name="db_path" label="File path" fullWidth onChange={handleInputChange("Path")} InputProps={{ className: classes.fieldColor }} value={global.Database.Path} />
        </div>;
    } else if (global.Database.Dbms === 'pgsql' || global.Database.Dbms === 'mysql') {
        fields = <div>
            <div style={{display: 'flex'}}>
                <TextField id="db_host" name="db_host" label="Hostname" onChange={handleInputChange("Hostname")} value={ global.Database.Hostname } InputProps={{ className: classes.fieldColor }} />
                <TextField id="db_port" name="db_port" label="Port" type="number" inputProps={{ min: "0" }} onChange={handleInputChange("Port")} value={ global.Database.Port } InputProps={{ className: classes.fieldColor }} />
            </div>
            <TextField id="db_username" name="db_username" label="Username" fullWidth onChange={handleInputChange("Username")} value={global.Database.Username} InputProps={{ className: classes.fieldColor }} />
            <TextField id="db_password" name="db_password" label="Password" fullWidth onChange={handleInputChange("Password")} value={global.Database.Password} InputProps={{ className: classes.fieldColor }} type="password" />
            <TextField id="db_database" name="db_database" label="Database" fullWidth onChange={handleInputChange("Database")} value={global.Database.Database} InputProps={{ className: classes.fieldColor }} />
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
                    <Select labelId="SelectDBMS" id="dbms" name="dbms" value={global.Database.Dbms} onChange={handleInputChange("Dbms")} classes={{ root: classes.fieldColor, icon: classes.fieldColor }}>
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
            <MuiAlert elevation={6} variant="filled" severity="error"> These settings are invalid! </MuiAlert>
        </Snackbar>
    </div>;
}