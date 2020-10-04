import { Login, Register } from "@/api/login";

export default {
    state: {
        Token: localStorage.getItem('scinna-token') || 'NO-TOKEN',
    },
    mutations: {
        loggedIn(state, payload) {
            Object.assign(state, payload);
        }
    },
    actions: {
        setToken(ctx, token) {
            localStorage.setItem('scinna-token', token);
        },
        removeToken() {
            localStorage.removeItem('scinna-token');
        },

        login(ctx, payload) {
            Login(ctx, payload)
        },

        register(ctx, payload) {
            Register(ctx, payload)
        },
    },
    modules: {

    }
}