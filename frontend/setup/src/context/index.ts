import React from 'react';

export const InitialState = {
    Database: {
        Dbms: 'pgsql',
        Path: '',
        Hostname: 'localhost',
        Port: 5432,
        Username: '',
        Password: '',
        Database: '',
    },
    Smtp: {
        Enabled: true,
        Hostname: '',
        Port: 587,
        Username: '',
        Password: '',
        Sender: '',
        TestReceiver: '',
    },
    Scinna: {
        RegistrationAllowed: "private",
        HeaderIPField: "",
        RateLimiting: 100,
        PicturePath: "/home/scinna/pictures/",
        WebURL: window.location.protocol + "//" + window.location.hostname,
    },
    User: {
        Username: "",
        Email: "",
        Password: "",
    }
}

export const AppContext = React.createContext(InitialState);
export const useStateValue = () => React.useContext(AppContext);