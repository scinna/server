import {Mutations} from "@/store/Mutations";
import {FetchUserInfos} from "@/api/User";
import {User} from "@/types/User";
import {AddInterceptor, RemoveInterceptor} from "@/api/Interceptors";
import {Commit} from "vuex";

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
    mutations: {
        [Mutations.LOAD_USER_TOKEN]: (state: AccountStateProps, token: string|null) => {
            state.Token = token;
        },
        [Mutations.LOGIN_RESPONSE]: (state: AccountStateProps, payload: {token: string} & User) => {
            state.User = { ...payload};
        },
        [Mutations.LOGOUT]: (state: AccountStateProps) => {
            state.Token = null;
            state.User = null;
        }
    },
    actions: {
        [Mutations.LOAD_USER_TOKEN]: ({commit}: { commit: Commit }) => new Promise((resolve, reject) => {
            const token = localStorage.getItem(TOKEN_KEY) || '';
            if (token.length > 0) {
                commit(Mutations.LOAD_USER_TOKEN, token);
                AddInterceptor(token);

                FetchUserInfos()
                    .then(resp => {
                        commit(Mutations.LOGIN_RESPONSE, resp.data as User);
                        resolve();
                    })
                    .catch(err => {
                        // If the server does not answer 401 it means that the token could be still valid and there is an error server side
                        // e.g. 502 if the server is down
                        console.log(err.response.status);
                        if (err.response.status === 401) {
                            console.log("removing");
                            localStorage.removeItem(TOKEN_KEY);
                            commit(Mutations.LOAD_USER_TOKEN, null);
                        }

                        reject();
                    })
            }
        }),
        [Mutations.LOGIN_RESPONSE]: ({commit}: { commit: Commit }, payload: { Token: string } & User) => new Promise((resolve) => {
            localStorage.setItem(TOKEN_KEY, payload.Token);

            commit(Mutations.LOGIN_RESPONSE, payload)
            resolve();
        }),
        [Mutations.LOGOUT]: ({commit}: { commit: Commit }) => new Promise((resolve) => {
            // @TODO send request to server to disable the token

            commit(Mutations.LOGOUT);
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
