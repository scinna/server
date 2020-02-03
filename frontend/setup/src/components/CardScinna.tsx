import React from 'react';
import {Link} from 'react-router-dom';
import { useForm } from 'react-hook-form';

import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';


const useStyles = makeStyles(theme => ({
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
    },
}));

export default function() {
    const classes = useStyles();
    const { register, handleSubmit, errors } = useForm();
    const onSubmit = (data: any) => console.log(data);
    console.log(errors);

    const [registration, setRegistration] = React.useState('public');
    
    const handleChange = (event: any) => {
        setRegistration(event.target.value);
    };

    return <div className="card above">
        <h4>About this server</h4>
        <form onSubmit={handleSubmit(onSubmit)}>
            <div className="content centered-form">
                <p>This is really important. Please <a href="https://github.com/scinna/server/wiki/First-launch#scinna-settings">follow the docs</a> to understand each options.</p>
                <FormControl className={classes.formControl} fullWidth>
                    <InputLabel id="registration">Server registration</InputLabel>
                    <Select labelId="registration" id="registration" value={registration} onChange={handleChange} inputRef={register({required: true})}>
                        <MenuItem value={"private"}>Private</MenuItem>
                        <MenuItem value={"public"}>Public</MenuItem>
                    </Select>
                </FormControl>
                <TextField id="scinna_header" label="IP Header" fullWidth inputRef={register({required: true, min: 1})}/>
                <TextField id="scinna_rate_limit" label="Rate limiting" fullWidth inputRef={register({required: true, min: 1})}/>
                <TextField id="scina_path" label="Picture path" fullWidth inputRef={register({required: true, min: 1})}/>
                <TextField id="scina_url" label="Web URL" fullWidth inputRef={register({required: true, min: 1})}/>
            </div>
            <div className="footer">
                <Link className="btn" to="/smtp">Back</Link>
                <Link className="btn" to="/user">Next</Link>
            </div>
        </form>
    </div>;
}