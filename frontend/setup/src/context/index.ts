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
    },
    Scinna: {
        Registration: "private",
        IPHeader: "",
        RateLimiting: "100",
        PicturePath: "/home/scinna/pictures/",
        WebURL: "https://i.example.com/",
    },
    User: {
        Username: "",
        Email: "",
        Password: "",
    }
}

export const AppContext = React.createContext(InitialState);
export const useStateValue = () => React.useContext(AppContext);