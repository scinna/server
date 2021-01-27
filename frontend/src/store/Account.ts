import {Mutations} from "@/store/Mutations";
import {FetchUserInfos} from "@/api/User";
import {User} from "@/types/User";

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
        })
    },
    getters: {
        isLoggedIn: (state: AccountStateProps) => state.Token && state.Token.length > 0 && state.User,
    }
};

export default accounts;
