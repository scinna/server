import React, {useState}                       from "react";
import {InputLabel, TextField, useFormControl} from "@material-ui/core";
import {Controller, useForm}                   from "react-hook-form";
import i18n                                    from "i18n-js";

type StatusType = {status: 'none'} | {status:'pending'} | {status: 'success' | 'error', message: string}

type IFormInputs = {
    password: string;
    repeat_password: string;
}

export function PasswordReset() {
    const [status, setStatus] = useState<StatusType>({status: 'none'});
    const {control, handleSubmit, reset} = useForm<IFormInputs>();

    const onSubmit = async (data: IFormInputs) => {

    }

    return <form onSubmit={handleSubmit(onSubmit)}>
        <InputLabel htmlFor="password">{i18n.t('forgotten_password.new_password')}</InputLabel>
        <Controller
            name="password"
            control={control}
            defaultValue=""
            render={({onChange, value}) => (
                <TextField
                    type="password"
                    value={value}
                    onChange={onChange}
                    disabled={status.status === 'pending'}
                    required
                    fullWidth
                />
            )}
        />
        <InputLabel htmlFor="repeat_password">{i18n.t('forgotten_password.repeat_password')}</InputLabel>
        <Controller
            name="repeat_password"
            control={control}
            defaultValue=""
            render={({onChange, value}) => (
                <TextField
                    type="password"
                    value={value}
                    onChange={onChange}
                    disabled={status.status === 'pending'}
                    required
                    fullWidth
                />
            )}
        />

        {
            status.status === 'success'
            &&
            <span>{status.message}</span>
        }
        {
            status.status === 'error'
            &&
            <span>{status.message}</span>
        }

    </form>
}