import axios from 'axios';
import { actionGotConfig } from '../actions/MainActions';

export function APIConfig(dispatch: any) {
    axios.get('/config')
        .then( (resp) => {
            dispatch(actionGotConfig(resp.data))
        })
}