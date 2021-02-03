import {Mutations} from "@/store/Mutations";
import {FetchUserInfos} from "@/api/User";
import {User} from "@/types/User";
import {AddInterceptor, RemoveInterceptor} from "@/api/Interceptors";

export const TOKEN_KEY = 'SCINNA_TOKEN';

interface AccountStateProps {
    User: User | null;
    Token: string | null;
}

const defaultState: AccountStateProps = {
    User: null,
    Token: null,
}

const accounts = {
    state: defaultState,
    mutations: {},
    actions: {
        [Mutations.LOAD_USER_TOKEN]: ({state}: { state: AccountStateProps}) => new Promise((resolve, reject) => {
            const token = localStorage.getItem(TOKEN_KEY) || '';
            if (token.length > 0) {
                state.Token = token;
                AddInterceptor(token);

                FetchUserInfos()
                    .then(resp => {
                        state.User = resp.data as User;
                        resolve();
                    })
                    .catch(err => {
                        // If the server does not answer 401 it means that the token could be still valid and there is an error server side
                        // e.g. 502 if the server is down
                        if (err.response.statusCode === 401) {
                            localStorage.removeItem(TOKEN_KEY);
                            state.Token = null;
                        }

                        reject();
                    })
            }
        }),
        [Mutations.LOGIN_RESPONSE]: ({state}: {state: AccountStateProps}, payload: {Token: string} & User) => new Promise((resolve, reject) => {
                state.Token = payload.Token;
                localStorage.setItem(TOKEN_KEY, payload.Token);

                state.User = payload;
                resolve();
        }),
        [Mutations.LOGOUT]: ({state}: {state: AccountStateProps}) => new Promise((resolve) => {
            // @TODO

            state.Token = null;
            state.User = null;
            localStorage.removeItem(TOKEN_KEY);
            RemoveInterceptor();

            resolve();
        })
    },
    getters: {
        isLoggedIn: (state: AccountStateProps) => state.Token && state.Token.length > 0 && state.User,
    }
};

export default accounts;
