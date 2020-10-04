import {API_URL} from "@/api/const";
import Vue from 'vue';

export const Login = ({dispatch, commit}, payload) => {
    Vue.axios.post(`${API_URL}auth/login`, payload.data)
        .then(r => {
            dispatch('setToken', r.data.Token);
            commit('loggedIn', r.data).then(() => {
                payload.router.push('/');
            });
        })
        .catch((e) => {
            if (e.response && e.response.status === 401) {
                commit('setErrorMessage', e.response.data.Message)
            } else {
                console.log(e);
            }
        })
}

export const Register = ({commit}, data) => {
    Vue.axios.post(`${API_URL}auth/register`, data)
        .then(r => {
            commit('registered', r.data)
        })
        .catch((e) => {
            console.log(e)
        })
}