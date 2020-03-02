import React from 'react';

export const InitialState = {
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
        Username: '',
        Token: '',
    }
}

export const AppContext = React.createContext(InitialState);
export const useStateValue = () => React.useContext(AppContext);