import React, {createContext, ReactNode, useContext, useState} from "react";
import {Token}                                                 from "../types/Token";
import {apiCall}                                               from "../utils/useApi";
import {useToken}                                              from "./TokenProvider";
import {isScinnaError, ScinnaError}                            from "../types/Error";
import useAsyncEffect                                          from "use-async-effect";

type Props = {
    children: ReactNode;
};

type TokenListProps = {
    loaded: boolean;
    status: null | 'pending' | 'success' | 'error';
    error?: string | null;
    tokens?: Token[];
}

type TokenListContextProps = TokenListProps & {
    revoke: (revokedToken: string) => void;
    refresh: () => void;
}

const defaultState: TokenListProps = {
    loaded: false,
    status: null,
    error: null,
    tokens: [],
};

const TokenListContext = createContext<TokenListContextProps>({
    ...defaultState,
    revoke: (Token) => {
    },
    refresh: () => {
    },
});

export default function AccountTokenProvider({children}: Props) {
    const {token} = useToken();
    const [context, setContext] = useState<TokenListProps>(defaultState);

    const refresh = async () => {
        const resp = await apiCall<Token[]>(token, {
            url: '/api/account/tokens'
        })

        if (isScinnaError(resp)) {
            setContext({...context, loaded: true, status: 'error', error: (resp as ScinnaError).Message})
            return;
        }

        setContext({...context, loaded: true, status: 'success', tokens: resp as Token[]})
    }

    const revoke = async (revokedToken: string) => {
        await apiCall(token, {
            url: '/api/account/tokens/' + revokedToken,
            method: 'DELETE',
        });

        await refresh();
    }

    useAsyncEffect(async () => {
        if (!context.loaded) {
            await refresh();
        }
    }, [context.loaded])

    return <TokenListContext.Provider value={{
        ...context,
        refresh,
        revoke
    }}>
        {children}
    </TokenListContext.Provider>
}

export function useAccountTokens(): TokenListContextProps {
    return useContext<TokenListContextProps>(TokenListContext);
}