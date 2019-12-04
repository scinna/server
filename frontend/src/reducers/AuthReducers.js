import {
    READ_TOKEN,
    DELETE_TOKEN,
}                                  from '../actions/AuthActions'
import { NO_TOKEN_AVAILABLE }      from '../middleware/LocalStorageMiddleware';;

const initialState = { token: '' };

export default function (state = initialState, action) {
    switch(action.type) {
        case READ_TOKEN:
            if (NO_TOKEN_AVAILABLE === action.payload.token || undefined === action.payload.token || null === action.payload.token || 0 >= action.payload.token.length) {
                return { ...state, token: action.payload.token };
            }
            return { ...state, token: NO_TOKEN_AVAILABLE };

        case DELETE_TOKEN:
            return { ...state, token: null, user: null };

        default:
            return { ...state };
    }
}