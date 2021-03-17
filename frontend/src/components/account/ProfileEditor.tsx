import {useToken} from "../../utils/TokenProvider";
import React, {useState} from "react";
import {ValidationErrors} from "../../types/ValidationErrors";
import {Controller, useForm} from "react-hook-form";
import i18n from "i18n-js";
import {Button, InputLabel, TextField} from "@material-ui/core";

import styles from '../../assets/scss/Profile.module.scss';

interface IFormInputs {
    Email: string;
    Password: string;
    Password2: string;
}

export function ProfileEditor() {
    const {token, userInfos} = useToken();
    const [validationErrors, setValidationErrors] = useState<ValidationErrors>();
    const [status, setStatus] = useState<null | 'error' | 'success' | 'pending'>(null);
    const [message, setMessage] = useState<null | String>(null);
    const {control, handleSubmit} = useForm<IFormInputs>();

    const onSubmit = async (data: IFormInputs) => {
        setValidationErrors(undefined);

        if (data.Password !== data.Password2) {
            setStatus('error');
            setValidationErrors({Violations: {'Password': i18n.t('errors.passwordNotMatching')}})

            return;
        }

        setStatus('pending');
        setMessage("");

        const response = await fetch("/api/account", {
            method: 'PUT',
            headers: {Authorization: 'Bearer ' + token},
            body: JSON.stringify(data),
        });

        if (!response.ok) {
            setStatus('error');
            try {
                const responseData = await response.json();
                if (responseData.Message) {
                    setMessage(responseData.Message);
                } else if (responseData.Violations) {
                    setValidationErrors(responseData);
                } else {
                    setMessage(i18n.t('errors.unknown'))
                }
            } catch {
                setMessage(i18n.t('errors.unknown'))
            }

            return
        }

        setStatus('success');
        setMessage(i18n.t('my_profile.account.success'));

    };

    // @TODO: Fix, we can't use defaultValue since a F5 on the page will first show this then pull the user
    return <div className={styles.TabProfile}>
        <h1>{i18n.t('my_profile.account.tab_name')}</h1>
        <form onSubmit={handleSubmit(onSubmit)}>
            <InputLabel htmlFor="Username">{i18n.t('registration.username')}</InputLabel>
            <Controller
                name="Username"
                control={control}
                defaultValue={userInfos?.Name}
                render={({onChange, value}) => <TextField onChange={onChange}
                                                          value={value}
                                                          error={validationErrors?.Violations?.hasOwnProperty("Username")}
                                                          helperText={validationErrors?.Violations?.Username}
                                                          disabled
                                                          fullWidth/>}
            />
            <InputLabel htmlFor="Email">{i18n.t('registration.email')}</InputLabel>
            <Controller
                name="Email"
                control={control}
                defaultValue={userInfos?.Email}
                render={({onChange, value}) => <TextField onChange={onChange}
                                                          value={value}
                                                          disabled={status === 'pending'}
                                                          error={validationErrors?.Violations?.hasOwnProperty("Email")}
                                                          helperText={validationErrors?.Violations?.Email}
                                                          required
                                                          fullWidth/>}
            />
            <InputLabel htmlFor="Password">{i18n.t('my_profile.account.current_password')}</InputLabel>
            <Controller
                name="CurrentPassword"
                control={control}
                defaultValue=""
                render={({onChange, value}) => <TextField onChange={onChange}
                                                          value={value}
                                                          disabled={status === 'pending'}
                                                          type="password"
                                                          error={validationErrors?.Violations?.hasOwnProperty("CurrentPassword")}
                                                          helperText={validationErrors?.Violations?.CurrentPassword}
                                                          required
                                                          fullWidth/>}
            />
            <InputLabel htmlFor="Password">{i18n.t('my_profile.account.new_password')}</InputLabel>
            <Controller
                name="Password"
                control={control}
                defaultValue=""
                render={({onChange, value}) => <TextField onChange={onChange}
                                                          value={value}
                                                          disabled={status === 'pending'}
                                                          type="password"
                                                          error={validationErrors?.Violations?.hasOwnProperty("Password")}
                                                          helperText={validationErrors?.Violations?.Password}
                                                          autoComplete="new-password"
                                                          fullWidth/>}
            />

            <InputLabel htmlFor="Password2">{i18n.t('registration.repeat_password')}</InputLabel>
            <Controller
                name="Password2"
                control={control}
                defaultValue=""
                render={({onChange, value}) => <TextField onChange={onChange}
                                                          disabled={status === 'pending'}
                                                          value={value}
                                                          type="password"
                                                          autoComplete="new-password"
                                                          fullWidth/>}
            />

            {
                message !== null
                &&
                <span className={
                    status === 'success'
                        ? styles.TabProfile__Message__Success
                        :
                    status === 'error'
                        ? styles.TabProfile__Message__Error
                        : ''
                }>{message}</span>
            }

            <Button type="submit" disabled={status === 'pending'}>
                {i18n.t('my_profile.account.update')}
            </Button>
        </form>
    </div>;
}