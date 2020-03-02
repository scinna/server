import IActionType from '../actions';
import { ACTION_MENU_STATE, ACTION_GOT_CONFIG } from '../actions/MainActions';

export default (state: any, action: IActionType) => {
    switch(action.type) {

        case ACTION_MENU_STATE:
            return { ...state, Main: { ...state.Main, menuOpened: action.payload.state } };

        case ACTION_GOT_CONFIG:
            return { ...state, Config: action.payload.config }

        default:
            return state;
    }
}