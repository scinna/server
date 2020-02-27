import {ACTION_UPDATE_DATABASE} from '../actions/database';
import {ACTION_UPDATE_SMTP} from '../actions/smtp';
import IActionType from '../actions';

export default (state: any, action: IActionType) => {
    switch (action.type) {
        case ACTION_UPDATE_DATABASE:
            return { ...state, Database: { ...state.Database, ...action.payload }}
        
        // @TODO: Go into his own reducer
        case ACTION_UPDATE_SMTP:
            return { ...state, Smtp: { ...state.Smtp, ...action.payload }}

        default:
            return state;
    }
}