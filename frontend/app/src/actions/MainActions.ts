import IActionType from './index';

export const ACTION_MENU_STATE = "MENU_SET_STATE";

export const actionMenuToggle = (state: boolean): IActionType => {
    return { type: ACTION_MENU_STATE, payload: { state } };
}