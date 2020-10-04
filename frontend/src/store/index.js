import Vue from 'vue'; import Vuex from 'vuex';

import User from './user';
import Server from './server';

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        ErrorMessage: "",
    },
    mutations: {
        clearError(state) {
            state.ErrorMessage = '';
        },
        setErrorMessage(state, payload) {
            state.ErrorMessage = payload;
        }
    },
    modules: {
        Server,
        User,
    }
})
