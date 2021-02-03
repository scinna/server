import axios from 'axios';

let tokenInterceptor = -1;

export function AddInterceptor(token: string) {
    tokenInterceptor = axios.interceptors.request.use(
        (config) => {
            config.headers["Authorization"] = "Bearer " + token;
            return config;
        },
        (err) => {
            return err;
        }
    )
}

export function RemoveInterceptor() {
    if (tokenInterceptor !== -1) {
        axios.interceptors.request.eject(tokenInterceptor);
        tokenInterceptor = -1;
    }
}
