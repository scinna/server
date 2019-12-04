import { readTokenAction, DELETE_TOKEN, GET_TOKEN, SET_TOKEN } from '../actions/AuthActions';

export const NO_TOKEN_AVAILABLE = "NO_TOKEN_AVAILABLE";

export default store => next => {
    return action => {

        switch (action.type) {
            case SET_TOKEN:
                localStorage.setItem('scinna_token', action.payload.token);
                break;
            case GET_TOKEN:
                let token = localStorage.getItem('scinna_token');
                if (null !== token) {
                    store.dispatch(readTokenAction({ token }));
                } else {
                    store.dispatch(readTokenAction({ token: NO_TOKEN_AVAILABLE, refreshToken: NO_TOKEN_AVAILABLE }));
                }
                break;
            case DELETE_TOKEN:
                localStorage.removeItem('scinna_token');
                store.dispatch(readTokenAction({ token: NO_TOKEN_AVAILABLE }));
                break;

            default:
                break;
        }

        return next(action);
    };
};