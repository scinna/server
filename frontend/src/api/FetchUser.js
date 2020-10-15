import axios from 'axios';
import {HOSTNAME} from "./constants";

export default function(commit, token) {
    axios.get(HOSTNAME + "/account/", { headers: { Authorization: "Brearer " + token }})
        .then(e => {
            console.log(e.status, e.statusCode, e.data)
        })
        .catch(e => {
            if (e.response.status === 401) {
                commit('removeToken')
            }
        })
}