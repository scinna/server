import {createContext, ReactNode, useContext, useState} from "react";
import {Token} from "../types/Token";
import {apiCall} from "../utils/useApi";
import {useToken} from "./TokenProvider";
import {isScinnaError, ScinnaError} from "../types/Error";

type Props = {
    children: ReactNode;
};

type TokenListProps = {
    loaded: boolean;
    error: string|null;
    tokens: Token[];
}

type TokenListContextProps = TokenListProps & {
    init: () => void;
    revoke: (token: Token) => void;
    refresh: () => void;
}

const defaultState: TokenListProps = {
    loaded: false,
    error:  null,
    tokens: [],
};

const TokenListContext = createContext<TokenListContextProps>({
    ...defaultState,
    init: () => {},
    revoke: (Token) => {},
    refresh: () => {},
});

export default function AccountTokenProvider({children}: Props) {
    const {token} = useToken();
    const [context, setContext] = useState<TokenListProps>(defaultState);

    const refresh = async function refresh() {
        const resp = await apiCall<Token[]>(token, {
            url: '/api/account/tokens'
        })

        if (isScinnaError(resp)) {
            setContext({...context, loaded: true, error: (resp as ScinnaError).Message})
            return;
        }

        setContext({...context, loaded: true, tokens: resp as Token[]})
    }

    const init = async () => await refresh();

    const revoke = async (token: Token) => {}

    return <TokenListContext.Provider value={{
        ...context,
        init,
        refresh,
        revoke
    }}>
        {children}
    </TokenListContext.Provider>
}

export function useAccountTokens(): TokenListContextProps {
    return useContext<TokenListContextProps>(TokenListContext);
}