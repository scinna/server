export const READ_TOKEN        = "READ_TOKEN";
export const GET_TOKEN         = "GET_TOKEN";
export const SET_TOKEN         = "SET_TOKEN";
export const DELETE_TOKEN      = "DELETE_TOKEN";

export const setTokenAction = (payload = {token: ''}) => {
    return {type: SET_TOKEN, payload};
};

export const getTokenAction = () => {
    return {type: GET_TOKEN};
};

export const readTokenAction = (payload = {token: ''}) => {
    return {type: READ_TOKEN, payload};
};

export const deleteTokenAction = () => {
    return {type: DELETE_TOKEN};
};
