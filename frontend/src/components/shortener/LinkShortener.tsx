import React from 'react';
import {Button, InputLabel, TextField} from "@material-ui/core";
import {useState} from "react";
import {useForm, Controller} from "react-hook-form";
import i18n from "i18n-js";

import styles from '../../assets/scss/linkshortener/shortener.module.scss';
import {apiCall} from "../../utils/useApi";
import {ShortenLink} from "../../types/ShortenLink";
import {useToken} from "../../context/TokenProvider";
import {isScinnaError} from "../../types/Error";
import {useShortenLink} from "../../context/ShortenLinkProvider";

type IFormInputs = {
    url: string;
}

export function LinkShortener() {
    const {token} = useToken();
    const [status, setStatus] = useState<null | 'error' | 'success' | 'pending'>(null);
    const {control, handleSubmit, reset} = useForm<IFormInputs>();
    const {refresh} = useShortenLink();

    const onSubmit = async (data: IFormInputs) => {
        setStatus('pending');

        const fd = new URLSearchParams();
        fd.append('url', data.url);

        const resp = await apiCall<ShortenLink>(token, {
            url: '/api/upload/shorten',
            method: 'POST',
            data: fd,
            mustNotBeStringified: true,
        });

        if (isScinnaError(resp)) {
            await setStatus('error');
            return
        }

        await setStatus('success');
        await reset();
        await refresh();
    };

    return <form className={styles.Shortener} onSubmit={handleSubmit(onSubmit)}>
        <InputLabel htmlFor="url">{i18n.t('shortener.link')}</InputLabel>
        <Controller
            name="url"
            control={control}
            defaultValue=""
            render={({onChange, value}) => (<TextField
                value={value}
                onChange={onChange}
                disabled={status === 'pending'}
                required
                fullWidth
            />)}
        />

        <Button type="submit" disabled={status === 'pending'}>
            {i18n.t('shortener.send')}
        </Button>
    </form>;
}