import {Button, InputLabel, TextField} from "@material-ui/core";
import React, {useEffect, useState} from "react";
import {Controller, useForm} from "react-hook-form";
import {useServerConfig} from "../utils/ServerConfigProvider";
import useAsyncEffect from "use-async-effect";

interface IFormInputs {
    Username: string;
    Email: string;
    Password: string;
    Password2: string;
    InviteCode: string;
}

export function Register() {
    const {Config} = useServerConfig();

    const {control, handleSubmit} = useForm<IFormInputs>();
    const [status, setStatus] = useState<null | 'error' | 'success' | 'pending'>(null);
    const [message, setMessage] = useState<null|String>(null);

    const onSubmit = async (data: IFormInputs) => {
        setStatus('pending');
        const response = await fetch('/api/auth/register', {
            method: 'POST',
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            console.log("err: ", response.status);
            setStatus('error');
            setMessage("Error occurred");
        } else {
            setStatus('success');

            const result = await response.json();
            setMessage(result.Message);
        }
    }

    return <form onSubmit={handleSubmit(onSubmit)}>
        <InputLabel htmlFor="Username">Username</InputLabel>
        <Controller
            name="Username"
            control={control}
            defaultValue=""
            render={({onChange, value}) => <TextField onChange={onChange}
                                                      value={value}
                                                      disabled={status === 'pending'}
                                                      required
                                                      fullWidth/>}
        />

        <InputLabel htmlFor="Email">Email</InputLabel>
        <Controller
            name="Email"
            control={control}
            defaultValue=""
            render={({onChange, value}) => <TextField onChange={onChange}
                                                      disabled={status === 'pending'}
                                                      value={value}
                                                      type="email"
                                                      required
                                                      fullWidth/>}
        />

        <InputLabel htmlFor="Password">Password</InputLabel>
        <Controller
            name="Password"
            control={control}
            defaultValue=""
            render={({onChange, value}) => <TextField onChange={onChange}
                                                      disabled={status === 'pending'}
                                                      value={value}
                                                      type="password"
                                                      autoComplete="new-password"
                                                      required
                                                      fullWidth/>}
        />

        <InputLabel htmlFor="Password2">Re-type your password</InputLabel>
        <Controller
            name="Password2"
            control={control}
            defaultValue=""
            render={({onChange, value}) => <TextField onChange={onChange}
                                                      disabled={status === 'pending'}
                                                      value={value}
                                                      type="password"
                                                      autoComplete="new-password"
                                                      required
                                                      fullWidth/>}
        />

        {
            !Config.RegistrationAllowed
            &&
            <>
                <InputLabel htmlFor="inviteCode">Invite code</InputLabel>
                <Controller
                    name="InviteCode"
                    control={control}
                    defaultValue=""
                    render={({onChange, value}) => <TextField onChange={onChange}
                                                              disabled={status === 'pending'}
                                                              value={value}
                                                              type="text"
                                                              autoComplete="new-password"
                                                              required
                                                              fullWidth/>}
                />
            </>
        }

        {
            message !== null
            &&
            <span className={status ?? ''}>{message}</span>
        }
        <Button type="submit" disabled={status === 'pending'}>Register</Button>
    </form>;
}