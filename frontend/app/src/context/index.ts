import React from 'react';

export const CtxInitialState = {
    Main: {
        menuOpened: false,
        registration: 'private'
    },
    Config: {
        Retreived: false,
        EmailAvailable: false,
        Registration: 'private',
    },
    User: {
        ID: 0,
        Username: '',
        Email: '',
        CreatedAt: new Date(),
        Role: 10, // Default to user, not admin
        Token: localStorage.getItem("scinna_token") ?? '',
    }
}

export const AppContext = React.createContext(CtxInitialState);
export const useStateValue = () => React.useContext(AppContext);