import {useToken}            from "../utils/TokenProvider";
import React, {useState}     from "react";
import {ValidationErrors}    from "../types/ValidationErrors";
import {Controller, useForm} from "react-hook-form";
import i18n                            from "i18n-js";
import {Button, InputLabel, TextField} from "@material-ui/core";

interface IFormInputs {
    Email: string;
    Password: string;
    Password2: string;
}

export function ProfileEditor() {
    const {userInfos} = useToken();
    const [validationErrors, setValidationErrors] = useState<ValidationErrors>();
    const [status, setStatus] = useState<null | 'error' | 'success' | 'pending'>(null);
    const [message, setMessage] = useState<null | String>(null);
    const {control, handleSubmit} = useForm<IFormInputs>();

    const onSubmit = async (data: IFormInputs) => {
        if (data.Password !== data.Password2) {
            setStatus('error');
            setValidationErrors({Violations: {'Password': i18n.t('errors.passwordNotMatching')}})

            return;
        }

        setStatus('pending');
        setMessage("");

    };

    return <div className="profileEditor">
        <h1>{i18n.t('my_profile.title')}</h1>
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
            <p>{i18n.t('my_profile.fill_if_changed')}</p>
            <InputLabel htmlFor="Password">{i18n.t('registration.password')}</InputLabel>
            <Controller
                name="Password"
                control={control}
                defaultValue=""
                render={({onChange, value}) => <TextField onChange={onChange}
                                                          value={value}
                                                          disabled={status === 'pending'}
                                                          error={validationErrors?.Violations?.hasOwnProperty("Password")}
                                                          helperText={validationErrors?.Violations?.Password}
                                                          autoComplete="new-password"
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
                                                          fullWidth/>}
            />

            {
                message !== null
                &&
                <span className={status ?? ''}>{message}</span>
            }

            <Button type="submit" disabled={status === 'pending'}>
                {i18n.t('my_profile.update')}
            </Button>
        </form>
    </div>;
}