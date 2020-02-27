import axios from 'axios';

export default (dbCreds: Object) => {
    return axios.post("/test/smtp", dbCreds)
}