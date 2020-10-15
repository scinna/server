import {createStore} from 'vuex'
import {HOSTNAME} from "@/api/constants";
import axios from "axios";

export default createStore({
    state: {
        User: null,
        Token: localStorage.getItem("scinna-token"),
        StatusMessage: '',
    },
    mutations: {
        removeToken(state) {
            localStorage.removeItem('scinna-token');
            state.Token = null;
        },
        setUser(state, user) {
            state.User = user;
            state.Token = user.Token;
        },
        error(state, err) {
            state.StatusMessage = err;
        }
    },
    actions: {
        login({commit}, data) {
            return new Promise(((resolve, reject) => {
                axios({url: HOSTNAME + '/auth/login', data, methods: 'POST'})
                    .then(r => {
                        commit('setUser', r.data)
                        localStorage.setItem('scinna-token', r.data.Token);
                    axios.defaults.headers.common['Authorization'] = "Bearer " + r.data.Token;
                        resolve(r.data);
                    })
                    .catch(e => {
                        commit('error', e.response.data.Message ?? 'Unknown error');
                        reject(e)
                    });
            }))
        },
        /** @TODO: Action logout, revoke the token **/
        logout({commit}) {
            commit('removeToken');
        }
    },
    getters: {
        isAuthenticated: state => !!state.Token,
        statusMessage: state => state.StatusMessage,
    },
    modules: {}
})
