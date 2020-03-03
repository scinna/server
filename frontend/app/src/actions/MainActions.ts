import IActionType from './index';
import { ILoginResponse } from '../api/models/Users';

export const ACTION_MENU_STATE = "MENU_SET_STATE";

export const actionMenuToggle = (state: boolean): IActionType => {
    return { type: ACTION_MENU_STATE, payload: { state } };
}

export interface IConfig {
    EmailAvailable: boolean,
    Registration: string
}

export const ACTION_GOT_CONFIG = "GOT_CONFIG";

export const actionGotConfig = (config: IConfig): IActionType => {
    return {type: ACTION_GOT_CONFIG, payload: { config: { Retreived: true, ...config }} }
}

export const ACTION_LOGGED_IN = "LOGGED_IN";

export const actionLoggedIn = (response: ILoginResponse): IActionType => {
    return { type: ACTION_LOGGED_IN, payload: { ...response } };
}

export const ACTION_BAD_TOKEN = "BAD_TOKEN";

export const actionBadToken = (): IActionType => {
    return { type: ACTION_BAD_TOKEN };
}

export const ACTION_LOG_OUT = "LOG_OUT";

export const actionLogout = (): IActionType => {
    return { type: ACTION_LOG_OUT };
}