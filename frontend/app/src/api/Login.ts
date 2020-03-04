import { scinnaxios, setAxiosToken } from './Axios';
import { actionLoggedIn, actionBadToken, actionLogout } from '../actions/MainActions';

interface ILoginData {
    Username: string,
    Password: string,
}

export function APIRegister(dispatch: any, email: boolean, data: ILoginData, actionAfter: any) {
    scinnaxios({method: 'POST', url: '/auth/register', data })
        .then( (resp) => {
            localStorage.setItem("scinna_token", resp.data.Token)
            setAxiosToken(resp.data.Token);
            actionAfter({ Severity: 'success', Message: 'Your account was created. Check your emails!' });
        })
        .catch( err => {
            switch (err.response.status) {
                case 400:
                    if (err.response.data.Errcode === 414) // Invalid Credentials
                        actionAfter({ Severity: 'warning', Message: 'Username or password invalid' });
                    else
                        actionAfter({ Severity: 'error', Message: err.response.data.Message });
                    break;
                case 500:
                    if (err.response.data.Errcode === 411)
                        actionAfter({ Severity: 'warning', Message: email ? 'Account created but there was an error sending the email. Contact the admin.' : 'Account created. Ask the admin to activate your account.' });
                    break;
                case 502:
                    actionAfter({ Severity: 'error', Message: 'Can\'t reach the server' });
                    break;
            }
        })
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

export function APILogout(dispatch: any, token: string) {
    scinnaxios({ method: 'GET', url: '/auth/logout'})
        .then(resp => {})
        .catch(err => {
            switch (err.response.status) {
                case 400:
                case 401:
                    let code = err.response.data.Errcode;
                    if (code === 399 || code === 400 || code === 401 || code === 403) { // Token not found, bad format, expired, ...
                        localStorage.removeItem("scinna_token")
                        dispatch(actionLogout());
                    }
                    break;
                case 410: // This is the normal response code
                    localStorage.removeItem("scinna_token")
                    dispatch(actionLogout());
                    break;
                case 502:
                    // @TODO: Add a global snackbar I think, instead of one in login
                    break;
            }
        })
}