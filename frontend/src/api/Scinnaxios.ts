import axios from 'axios';

export function AddInterceptor(token: string) {
    axios.interceptors.request.use(
        (config) => {
            config.headers["Authorization"] = "Bearer " + token;
            return config;
        },
        (err) => {
            return err;
        }
    )
}