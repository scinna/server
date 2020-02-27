import React from 'react';

export const InitialState = {
    Database: {
        Dbms: 'pgsql',
        Path: '',
        Hostname: '',
        Port: '',
        Username: '',
        Password: '',
        Database: '',
    },
    Smtp: {
        Enabled: true,
        Hostname: '',
        Port: '',
        Username: '',
        Password: '',
        Sender: '',
        TestReceiver: '',
    }
}

export const AppContext = React.createContext(InitialState);
export const useStateValue = () => React.useContext(AppContext);