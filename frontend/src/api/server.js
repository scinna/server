import {API_URL} from "@/api/const";
import Vue from 'vue';

export const FetchServerConfig = (commit) => {
    Vue.axios.get(`${API_URL}infos`)
        .then(r => {
            commit('gotServerConfig', r.data)
        })
        .catch((e) => {
            console.log(e)

            // @TODO: 'Crash' the app
            // Can't get the server config means the app won't work at all
        })
}