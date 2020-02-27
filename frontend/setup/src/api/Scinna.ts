import axios from 'axios';

export default (settings: Object) => {
    return axios.post("/scinna", settings)
}   