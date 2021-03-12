import {createContext, ReactNode, useContext, useEffect, useState} from "react";
import * as React                                                  from "react";

const LS_SCINNA_KEY = "SCINNA_TOKEN";

type UserInfos = null | {
    UserID: String,
    Name: String,
    Email: String,
    IsAdmin: Boolean,
}

type TokenProps = {
    token: string | null,
    loaded: Boolean,
    userInfos: UserInfos
};

type TokenContextProps = TokenProps & {
    init: () => void,
    setUserInfo: (token: string, userInfos: UserInfos) => void,
    isAuthenticated: () => Boolean;
};

const TokenContext = createContext<TokenContextProps>({
    token: null,
    loaded: false,
    init: () => {
    },
    isAuthenticated: () => false,
    setUserInfo: () => {},
    userInfos: null,
});

type Props = {
    children: ReactNode,
}

export default function TokenProvider({children}: Props) {
    const [context, setContext] = useState<TokenProps>({
        token: null,
        loaded: false,
        userInfos: null,
    });

    useEffect(() => {
        if (context.token) {
            localStorage.setItem(LS_SCINNA_KEY, context.token);
        } else {
            localStorage.removeItem(LS_SCINNA_KEY);
        }
    }, [context.token]);

    async function init(): Promise<void> {
        if (context.loaded) {
            return;
        }

        const token = localStorage.getItem(LS_SCINNA_KEY);
        let userInfos: UserInfos = null;
        if (token) {
            const response = await fetch("/api/account", {headers: {"Authorization": "Bearer " + token}})
            if (!response.ok) {
                localStorage.remove(LS_SCINNA_KEY);
                setContext({...context, loaded: true, token: null});

                return;
            }

            userInfos = await response.json();
        }
        setContext({...context, loaded: true, token, userInfos});
    }

    const setUserInfo = (token: string, userInfos: UserInfos) => {
        setContext({ ...context, token, userInfos, loaded: true})
    }

    const isAuthenticated = () => context.userInfos !== null;

    return (<TokenContext.Provider value={{
        ...context,
        init,
        isAuthenticated,
        setUserInfo,
    }}>
        {children}
    </TokenContext.Provider>)
}

export function useToken(): TokenContextProps {
    return useContext<TokenContextProps>(TokenContext);
}