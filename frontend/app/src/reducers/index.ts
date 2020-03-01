import IActionType from '../actions';
import { ACTION_MENU_STATE } from '../actions/MainActions';

export default (state: any, action: IActionType) => {
    switch(action.type) {

        case ACTION_MENU_STATE:
            return { ...state, Main: { ...state.Main, menuOpened: action.payload.state } };

        default:
            return state;
    }
}