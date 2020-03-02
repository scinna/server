import axios from 'axios';
import { actionGotConfig } from '../actions/MainActions';

export function getConfig(dispatch: any) {
    axios.get('/config')
        .then( (resp) => {
            dispatch(actionGotConfig(resp.data))
        })
}