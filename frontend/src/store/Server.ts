import {Mutations} from "@/store/Mutations";

export interface ServerProps {
    RegistrationAllowed: boolean;
    Validation: string | null;
    WebURL: string | null;
}

const defaultState: ServerProps = {
    RegistrationAllowed: false,
    Validation: null,
    WebURL: null,
}

const server = {
    state: defaultState,
    mutations: {
        [Mutations.GOT_SERVER_INFOS]: (state: ServerProps, payload: ServerProps) => {
            state.RegistrationAllowed = payload.RegistrationAllowed;
            state.Validation = payload.Validation;
            state.WebURL = payload.WebURL;
        }
    },
    actions: {},
};

export default server;
