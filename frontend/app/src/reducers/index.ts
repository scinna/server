import IActionType from '../actions';
import { CtxInitialState } from '../context';
import { ACTION_MENU_STATE, ACTION_GOT_CONFIG, ACTION_LOGGED_IN, ACTION_BAD_TOKEN } from '../actions/MainActions';

/**
 * @TODO: Find a way to split the context (Maybe have multiple context ? This will grow a lot the App file...)
 */
export default (state: any, action: IActionType) => {
    switch(action.type) {

        case ACTION_MENU_STATE:
            return { ...state, Main: { ...state.Main, menuOpened: action.payload.state } };

        case ACTION_GOT_CONFIG:
            return { ...state, Config: action.payload.config }

        case ACTION_LOGGED_IN:
            return { ...state, User: {
                ...action.payload.CurrentUser,
                Token: action.payload.Token
            }}

        case ACTION_BAD_TOKEN: 
            return { ...state, User: { ...CtxInitialState.User, Token: '' }}

        default:
            return state;
    }
}