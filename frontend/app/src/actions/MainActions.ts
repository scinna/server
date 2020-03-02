import IActionType from './index';

export const ACTION_MENU_STATE = "MENU_SET_STATE";

export const actionMenuToggle = (state: boolean): IActionType => {
    return { type: ACTION_MENU_STATE, payload: { state } };
}

export interface IConfig {
    EmailAvailable: boolean,
    Registration: string
}

export const ACTION_GOT_CONFIG = "GOT_CONFIG";

export const actionGotConfig = (config: IConfig) => {
    return {type: ACTION_GOT_CONFIG, payload: { config: { Retreived: true, ...config }} }
}