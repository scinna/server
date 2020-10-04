import {FetchServerConfig} from "@/api/server";

export default {
    state: {
        RegistrationAllowed: false,
        Validation: 'admin',
        WebURL: 'https://scinna.drx/',
    },
    mutations: {
        gotServerConfig(state, cfg) {
            Object.assign(state, cfg)
        }
    },
    actions: {
        fetchServerConfig(ctx) {
            FetchServerConfig(ctx.commit);
        }
    },
    modules: {

    }
}