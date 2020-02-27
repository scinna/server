import axios from 'axios';

export default (dbCreds: Object) => {
    return axios.post("/test/db", dbCreds)
}