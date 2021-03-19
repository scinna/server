import {Controller, useForm}           from "react-hook-form";
import {Button, InputLabel, TextField} from "@material-ui/core";
import React, {useState}               from "react";
import {Redirect}                      from "react-router-dom";
import i18n                            from "i18n-js";
import {useToken}                      from "../context/TokenProvider";

import '../assets/scss/Login.scss';

interface IFormInputs {
    Username: string;
    Password: string;
}

export function Login() {
    const [status, setStatus] = useState<null | 'error' | 'success' | 'pending'>(null);
    const [message, setMessage] = useState<null | String>(null);
    const {control, handleSubmit, reset} = useForm<IFormInputs>();

    const {isAuthenticated, setUserInfo} = useToken();

    // @TODO: Reset on failure
    const onSubmit = async function (data: IFormInputs) {
        setStatus('pending');
        setMessage('');

        const response = await fetch('/api/auth', {
            method: 'POST',
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            setStatus('error');
            reset({ ...data, Password: '' });
            try {
                const responseData = await response.json();
                if (responseData.Message) {
                    setMessage(responseData.Message);
                } else {
                    setMessage(i18n.t('errors.unknown'))
                }
            } catch {
                setMessage(i18n.t('errors.unknown'))
            }

            return;
        }

        setStatus('success');
        const responseData = await response.json();
        setUserInfo(responseData.Token, responseData);
    }

    if (isAuthenticated()) {
        return <Redirect to="/"/>;
    }

    return <div className="centeredBlock login">
        <h1>{i18n.t('login.title')}</h1>
        <form onSubmit={handleSubmit(onSubmit)}>
            <InputLabel htmlFor="Username">{i18n.t('login.username')}</InputLabel>
            <Controller
                name="Username"
                control={control}
                defaultValue=""
                render={({onChange, value}) => (
                    <TextField
                        value={value}
                        onChange={onChange}
                        disabled={status === 'pending'}
                        required
                        fullWidth
                    />
                )}
            />

            <InputLabel htmlFor="Password">{i18n.t('login.password')}</InputLabel>
            <Controller
                name="Password"
                control={control}
                defaultValue=""
                render={({onChange, value}) => (
                    <TextField
                        type="password"
                        value={value}
                        onChange={onChange}
                        disabled={status === 'pending'}
                        required
                        fullWidth
                    />
                )}
            />

            {
                message
                && <p className="error">{message}</p>
            }

            <Button type="submit" disabled={status === 'pending'}>
                {i18n.t('login.button')}
            </Button>
        </form>
    </div>;
}