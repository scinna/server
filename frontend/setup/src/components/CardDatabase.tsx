import React from 'react';
import {Link, Redirect} from 'react-router-dom';

import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';

import TestDB from '../api/Database';

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
    Dbms: 'pgsql',
    ButtonsEnabled: true,
};

export default function() {
    const classes = useStyles();
    /** 
     * @TODO: use Context to store the field inputs
     * so that changing page and going back keeps the values
     */
    
    const [state, setState] = React.useState(initialState);

    const [input, setInput] = React.useState({ dbms: state.Dbms })
    const handleInputChange = (e: any) => setInput({
        ...input,
        [e.currentTarget.name]: e.currentTarget.value
    })

    const handleChange = (event: any) => {
        setState({
            ...state,
            Dbms: event.target.value,
        });
        handleInputChange({ currentTarget: event.target})
    };

    const submit = (e: any) => {
        e.preventDefault();

        TestDB(input)
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
    if (state.Dbms === 'sqlite3') {
        fields = <div>
            <TextField id="db_path" name="db_path" label="File path" fullWidth onChange={handleInputChange} InputProps={{ className: classes.fieldColor }} />
        </div>;
    } else if (state.Dbms === 'pgsql' || state.Dbms === 'mysql') {
        fields = <div>
            <div style={{display: 'flex'}}>
                <TextField id="db_host" name="db_host" label="Hostname" onChange={handleInputChange} InputProps={{ className: classes.fieldColor }} />
                <TextField id="db_port" name="db_port" label="Port" type="number" inputProps={{ min: "0" }} onChange={handleInputChange} InputProps={{ className: classes.fieldColor }} />
            </div>
            <TextField id="db_username" name="db_username" label="Username" fullWidth onChange={handleInputChange} InputProps={{ className: classes.fieldColor }} />
            <TextField id="db_password" name="db_password" label="Password" fullWidth onChange={handleInputChange} InputProps={{ className: classes.fieldColor }} type="password" />
            <TextField id="db_database" name="db_database" label="Database" fullWidth onChange={handleInputChange} InputProps={{ className: classes.fieldColor }} />
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
                    <Select labelId="SelectDBMS" id="dbms" name="dbms" value={state.Dbms} onChange={handleChange} classes={{ root: classes.fieldColor, icon: classes.fieldColor }}>
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
    </div>;
}