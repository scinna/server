import IActionType from '../actions';

import {ACTION_UPDATE_DATABASE} from '../actions/database';
import {ACTION_UPDATE_SMTP} from '../actions/smtp';
import { ACTION_UPDATE_SCINNA } from '../actions/scinna';
import { ACTION_UPDATE_USER } from '../actions/user';

export default (state: any, action: IActionType) => {
    switch (action.type) {
        case ACTION_UPDATE_DATABASE:
            return { ...state, Database: { ...state.Database, ...action.payload }}
        
        // @TODO: Go into his own reducer
        case ACTION_UPDATE_SMTP:
            return { ...state, Smtp: { ...state.Smtp, ...action.payload }}

        // @TODO: Go into his own reducer
        case ACTION_UPDATE_SCINNA:
            return { ...state, Scinna: { ...state.Scinna, ...action.payload }}

        // @TODO: Go into his own reducer
        case ACTION_UPDATE_USER:
            return { ...state, User: { ...state.User, ...action.payload }}

        default:
            return state;
    }
}