import {Controller, useForm} from "react-hook-form";

import '../assets/scss/Login.scss';
import {Button, InputLabel, TextField} from "@material-ui/core";
import React, {useState} from "react";
import i18n from "i18n-js";

interface IFormInputs {
    Username: string;
    Password: string;
}

export function Login() {
    const [status, setStatus] = useState<null | 'error' | 'success' | 'pending'>(null);
    const {control, handleSubmit} = useForm<IFormInputs>();

    const onSubmit = async function (data: IFormInputs) {
        setStatus('pending');
    }

    return <div className="centeredBlock">
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

            <Button type="submit" disabled={status === 'pending'}>
                {i18n.t('login.button')}
            </Button>
        </form>
    </div>;
}