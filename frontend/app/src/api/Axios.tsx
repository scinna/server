import axios from 'axios';

import React from 'react';
import { Redirect } from 'react-router';

import {APIConfig} from './Config'; 
import {APICheckToken} from './Login';

export const scinnaxios = axios.create()

let CurrentInterceptor: number = -1;

export const setAxiosToken = (token: string) => {
    if (CurrentInterceptor !== null) {
        scinnaxios.interceptors.request.eject(CurrentInterceptor);
    }

    CurrentInterceptor = scinnaxios.interceptors.request.use((cfg) => {
        cfg.headers.authorization = "Bearer " + token;
        return cfg;
    }, (err) => Promise.reject(err));
}

export function AxiosMiddlishware(dispatch: Function): JSX.Element|null {
    const token = localStorage.getItem("scinna_token")
    
    if (!token || token.length === 0) {
        return <Redirect to="/"/>
    }

    if (CurrentInterceptor === -1) {
        APIConfig(dispatch);
        setAxiosToken(token);
        APICheckToken(dispatch, token);

    }

    return null
}