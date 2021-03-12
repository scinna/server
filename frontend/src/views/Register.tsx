import {Button, InputLabel, TextField} from "@material-ui/core";
import React, {useState} from "react";
import {Controller, useForm} from "react-hook-form";
import {useServerConfig} from "../utils/ServerConfigProvider";

import '../assets/scss/Register.scss';
import i18n from "i18n-js";
import {ValidationErrors} from "../types/ValidationErrors";

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
    const [validationErrors, setValidationErrors] = useState<ValidationErrors>();
    const [status, setStatus] = useState<null | 'error' | 'success' | 'pending'>(null);
    const [message, setMessage] = useState<null | String>(null);

    const onSubmit = async (data: IFormInputs) => {
        setStatus('pending');
        const response = await fetch('/api/auth/register', {
            method: 'POST',
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            setStatus('error');
            try {
                const data = await response.json();
                if (data.Violations) {
                    setValidationErrors(data.Violations);
                } else {
                    setMessage(data.Message);
                }
            } catch {
                setMessage("Error occurred");
            }
        } else {
            setStatus('success');

            const result = await response.json();
            setMessage(result.Message);
        }
    }

    return <div className="centeredBlock">
        <h1>{i18n.t('registration.title')}</h1>
        <form onSubmit={handleSubmit(onSubmit)}>
            <InputLabel htmlFor="Username">{i18n.t('registration.username')}</InputLabel>
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

            <InputLabel htmlFor="Email">{i18n.t('registration.email')}</InputLabel>
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

            <InputLabel htmlFor="Password">{i18n.t('registration.password')}</InputLabel>
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

            <InputLabel htmlFor="Password2">{i18n.t('registration.repeatPassword')}</InputLabel>
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
                    <InputLabel htmlFor="inviteCode">{i18n.t('registration.inviteCode')}</InputLabel>
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

            <Button type="submit" disabled={status === 'pending'}>
                {i18n.t('registration.button')}
            </Button>
        </form>
    </div>;
}