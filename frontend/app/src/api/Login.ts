import { scinnaxios, setAxiosToken } from './Axios';
import { actionLoggedIn, actionBadToken } from '../actions/MainActions';

interface ILoginData {
    Username: string,
    Password: string,
}

export function APILogin(dispatch: any, data: ILoginData, actionAfter: any) {
    scinnaxios({method: 'POST', url: '/auth/login', data })
        .then( (resp) => {
            localStorage.setItem("scinna_token", resp.data.Token)
            setAxiosToken(resp.data.Token);
            dispatch(actionLoggedIn(resp.data));
        })
        .catch( err => {
            switch (err.response.status) {
                case 400:
                    if (err.response.data.Errcode === 414) // Invalid Credentials
                        actionAfter({ Severity: 'warning', Message: 'Username or password invalid' });
                    else
                        actionAfter({ Severity: 'error', Message: err.response.data.Message });
                    break;
                case 502:
                    actionAfter({ Severity: 'error', Message: 'Can\'t reach the server' });
                    break;
            }
        })
}

export function APICheckToken(dispatch: any, token: string) {
    scinnaxios({ method: 'GET', url: '/auth/token'})
        .then(resp => {
            dispatch(actionLoggedIn({ CurrentUser: resp.data, Token: token}));
        })
        .catch(err => {
            console.log(err.response)
            switch (err.response.status) {
                case 400:
                case 401:
                    let code = err.response.data.Errcode;
                    if (code === 399 || code === 400 || code === 401 || code === 403) { // Token not found, bad format, expired, ...
                        localStorage.removeItem("scinna_token")
                        dispatch(actionBadToken());
                    }
                    break;
                case 502:
                    // @TODO: Add a global snackbar I think, instead of one in login
                    break;
            }
        })
}