import axios from 'axios';

export default (userAccount: Object) => {
    return axios.post("/user", userAccount)
}