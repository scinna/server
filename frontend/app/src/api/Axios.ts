import axios from 'axios';

export const scinnaxios = axios.create()

let CurrentInterceptor: number;

export const setAxiosToken = (token: string) => {
    if (CurrentInterceptor !== null) {
        scinnaxios.interceptors.request.eject(CurrentInterceptor);
    }

    CurrentInterceptor = scinnaxios.interceptors.request.use((cfg) => {
        cfg.headers.authorization = "Bearer " + token;
        return cfg;
    }, (err) => Promise.reject(err));
}