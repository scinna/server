import React, {createContext, ReactNode, useContext, useState} from "react";
import {apiCall} from "../utils/useApi";
import {useToken} from "./TokenProvider";
import {isScinnaError, ScinnaError} from "../types/Error";
import useAsyncEffect from "use-async-effect";
import {Collection} from "../types/Collection";

type Props = {
    children: ReactNode;
};

type BrowserProps = {
    loaded: boolean;
    pending: boolean;
    status: null | 'pending' | 'success' | 'error';
    error?: string | null;
    username?: string;
    path?: string;
    collection?: Collection;
    isMine: boolean;
}

type BrowserContextProps = BrowserProps & {
    browse: (username: string, path?: string) => void;
    refresh: () => void;
}

const defaultState: BrowserProps = {
    loaded: false,
    pending: true,
    status: null,
    error: null,
    isMine: false,
};

const TokenListContext = createContext<BrowserContextProps>({
    ...defaultState,
    browse: (username, path) => {},
    refresh: () => {
    },
});

export default function BrowserProvider({children}: Props) {
    const {token, userInfos} = useToken();
    const [context, setContext] = useState<BrowserProps>(defaultState);

    const refresh = async () => {
        await setContext({...context, pending: true})
        const response = await apiCall<Collection>(token, {
            url: '/api/browse/' + context.username + '/' + (context.path ?? ''),
            method: 'GET',
            canBeUnauthed: true,
        });

        if (isScinnaError(response)) {
            await setContext({...context, loaded: true, pending: false, error: (response as ScinnaError).Message, isMine: false})
            return
        }

        await setContext({...context, loaded: true, pending: false, collection: (response as Collection), isMine: context.username?.toLowerCase() === userInfos?.Name.toLowerCase()})
    }

    const browse = async (username: string, path?: string) => {
        await setContext({...context, username, path: path ?? ''})
    }

    useAsyncEffect(async () => {
        if (context.username)
            await refresh();
    }, [context.loaded, context.username, context.path])

    return <TokenListContext.Provider value={{
        ...context,
        browse,
        refresh,
    }}>
        {children}
    </TokenListContext.Provider>
}

export function useBrowser(): BrowserContextProps {
    return useContext<BrowserContextProps>(TokenListContext);
}