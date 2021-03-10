import React, {useContext}                  from 'react';
import {createContext, ReactNode, useState} from "react";
import useAsyncEffect                       from "use-async-effect";

type Props = {
    children: ReactNode,
}

type ServerConfig = { RegistrationAllowed: boolean; Validation: string; WebURL: string; }
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
        WebURL: ''
    },
};

const ServerConfigContext = createContext<ServerConfigContextProps>(defaultValues);

export default function ServerConfigProvider({children}: Props) {
    const [context, setContext] = useState<ServerConfigProps>(defaultValues);

    useAsyncEffect(async () => {
        if (!context.Loaded) {
            const response = await fetch('/api/server/infos');
            const data = await response.json();
            setContext({ Loaded: true, Config: data });
        }
    }, [context.Loaded])

    return (<ServerConfigContext.Provider value={{
        ...context
    }}>
        {children}
    </ServerConfigContext.Provider>)
}

export function useServerConfig(): ServerConfigContextProps {
    return useContext<ServerConfigContextProps>(ServerConfigContext);
}
