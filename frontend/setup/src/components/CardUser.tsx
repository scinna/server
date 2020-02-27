import React from 'react';
import {Link} from 'react-router-dom';
import { useForm } from 'react-hook-form';

import TextField from '@material-ui/core/TextField';

/**
 * When the user click on Next, the app should send the data to the /save endpoint to test them
 * If there is any error (DB settings invalid, SMTP settings invalid, Scinna settings invalid
 * Can't register the user), the app should not go to the finale card and display an error.
 */

export default function() {
    const { register, handleSubmit, errors } = useForm();
    const onSubmit = (data: any) => console.log(data);
        console.log(errors);

    return <div className="card above">
        <h4>Create your account</h4>
        <form onSubmit={handleSubmit(onSubmit)}>
            <div className="content">
                <p>Creating the admin account.</p>
                <TextField id="user_name" label="Username" fullWidth inputRef={register({required: true, min: 1})}/>
                <TextField id="user_mail" label="Email" fullWidth inputRef={register({required: true, min: 1})}/>
                <TextField id="user_pass" label="Password" type="password" fullWidth inputRef={register({required: true, min: 1})}/>
                <TextField id="user_pwd2" label="Repeat password" type="password" fullWidth inputRef={register({required: true, min: 1})}/>
            </div>
            <div className="footer">
                <Link className="btn" to="/scinna">Back</Link>
                <Link className="btn" to="/finale">Next</Link>
            </div>
        </form>
    </div>;
}