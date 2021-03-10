import {Button, InputLabel, TextField} from "@material-ui/core";
import React                           from "react";
import {Controller, useForm}           from "react-hook-form";
import {useServerConfig}               from "../utils/ServerConfigProvider";

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

    const onSubmit = (data: IFormInputs) => {
        console.log(data);
    }

    return <form onSubmit={handleSubmit(onSubmit)}>
        <InputLabel htmlFor="Username">Username</InputLabel>
        <Controller
            name="Username"
            control={control}
            defaultValue=""
            render={({onChange, value}) => <TextField onChange={onChange} value={value} fullWidth/>}
        />

        <InputLabel htmlFor="Email">Email</InputLabel>
        <Controller
            name="Email"
            control={control}
            defaultValue=""
            render={({onChange, value}) => <TextField onChange={onChange}
                                                      value={value}
                                                      type="email"
                                                      fullWidth/>}
        />

        <InputLabel htmlFor="Password">Password</InputLabel>
        <Controller
            name="Password"
            control={control}
            defaultValue=""
            render={({onChange, value}) => <TextField onChange={onChange}
                                                      value={value}
                                                      type="password"
                                                      autoComplete="new-password"
                                                      fullWidth/>}
        />

        <InputLabel htmlFor="Password2">Re-type your password</InputLabel>
        <Controller
            name="Password2"
            control={control}
            defaultValue=""
            render={({onChange, value}) => <TextField onChange={onChange}
                                                      value={value}
                                                      type="password"
                                                      autoComplete="new-password"
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
                                                              value={value}
                                                              type="text"
                                                              autoComplete="new-password"
                                                              fullWidth/>}
                />
            </>
        }

        <Button type="submit">Register</Button>
    </form>;
}