import axios from 'axios';

export default (smtpCreds: Object) => {
    return axios.post("/test/smtp", smtpCreds)
}