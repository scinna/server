import axios from 'axios';
import { GetUrl } from './Debug';

export default (dbCreds: Object) => {
    return axios.post(GetUrl("/test/db"), dbCreds)
}