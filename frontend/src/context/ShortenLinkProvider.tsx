import React, {createContext, ReactNode, useContext, useState} from "react";
import {apiCall} from "../utils/useApi";
import {useToken} from "./TokenProvider";
import {isScinnaError, ScinnaError} from "../types/Error";
import useAsyncEffect from "use-async-effect";
import {ShortenLink} from "../types/ShortenLink";

type Props = {
    children: ReactNode;
};

type ShortenLinkProps = {
    loaded: boolean;
    status: null | 'pending' | 'success' | 'error';
    error?: string | null;
    links?: ShortenLink[];
}

type ShortenLinkContextProps = ShortenLinkProps & {
    refresh: () => void;
}

const defaultState: ShortenLinkProps = {
    loaded: false,
    status: null,
    error: null,
    links: [],
};

const ShortenLinkContext = createContext<ShortenLinkContextProps>({
    ...defaultState,
    refresh: () => {},
});

export default function ShortenLinkProvider({children}: Props) {
    const {token} = useToken();
    const [context, setContext] = useState<ShortenLinkProps>(defaultState);

    const refresh = async () => {
        const resp = await apiCall<ShortenLink[]>(token, {
            url: '/api/account/shorten_links'
        })

        if (isScinnaError(resp)) {
            setContext({...context, loaded: true, status: 'error', error: (resp as ScinnaError).Message})
            return;
        }

        setContext({...context, loaded: true, status: 'success', links: resp as ShortenLink[]})
    }

    useAsyncEffect(async () => {
        await refresh();
    }, [context.loaded, token])

    return <ShortenLinkContext.Provider value={{
        ...context,
        refresh,
    }}>
        {children}
    </ShortenLinkContext.Provider>
}

export function useShortenLink(): ShortenLinkContextProps {
    return useContext<ShortenLinkContextProps>(ShortenLinkContext);
}