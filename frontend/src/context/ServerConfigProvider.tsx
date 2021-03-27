import React, {useContext}                  from 'react';
import {createContext, ReactNode, useState} from "react";
import useAsyncEffect                       from "use-async-effect";
import {useToken} from "./TokenProvider";

type Props = {
    children: ReactNode,
}

type ServerConfig = {
    RegistrationAllowed: boolean;
    Validation: string;
    WebURL: string;
    CustomBranding: string;

    ScinnaVersion: string;
}
type ServerConfigProps = {
    Loaded: boolean,
    Config: ServerConfig,
}

type ServerConfigContextProps = ServerConfigProps & {};

const defaultValues: ServerConfigContextProps = {
    Loaded: false,
    Config: {
        RegistrationAllowed: false,
        Validation: 'email',
        WebURL: '',
        CustomBranding: '',
        ScinnaVersion: '',
    },
};

const ServerConfigContext = createContext<ServerConfigContextProps>(defaultValues);

export default function ServerConfigProvider({children}: Props) {
    const [context, setContext] = useState<ServerConfigProps>(defaultValues);
    const { token } = useToken();

    useAsyncEffect(async () => {
        const response = await fetch('/api/server/infos',
            token ? { headers: { Authorization: 'Bearer ' + token} } : {}
        );
        const data = await response.json();
        setContext({ Loaded: true, Config: data });
    }, [token])

    return (<ServerConfigContext.Provider value={{
        ...context
    }}>
        {children}
    </ServerConfigContext.Provider>)
}

export function useServerConfig(): ServerConfigContextProps {
    return useContext<ServerConfigContextProps>(ServerConfigContext);
}
