import React, {createContext, ReactNode, useContext, useState} from "react";
import {Token}                                                 from "../types/Token";
import {apiCall}                                               from "../utils/useApi";
import {useToken}                                              from "./TokenProvider";
import {isScinnaError, ScinnaError}                            from "../types/Error";
import useAsyncEffect                                          from "use-async-effect";
import {InviteCode}                                            from "../components/server/InviteCode";

type Props = {
    children: ReactNode;
};

type InviteListProps = {
    loaded: boolean;
    status: null | 'pending' | 'success' | 'error';
    error?: string | null;
    invites?: InviteCode[];
    newlyGeneratedCode: string;
}

type InviteListContextProps = InviteListProps & {
    generate: () => void;
    remove: (code: string) => void;
    refresh: () => void;
}

const defaultState: InviteListProps = {
    loaded: false,
    status: null,
    error: null,
    invites: [],
    newlyGeneratedCode: '',
};

const InviteListContext = createContext<InviteListContextProps>({
    ...defaultState,
    generate: () => {
    },
    remove: (code: string) => {
    },
    refresh: () => {
    },
});

export default function InviteCodeProvider({children}: Props) {
    const {token} = useToken();
    const [context, setContext] = useState<InviteListProps>(defaultState);

    const refresh = async () => {
        setContext({...context, status: 'pending'})
        const resp = await apiCall<InviteCode[]>(token, {
            url: '/api/server/admin/invite'
        })

        if (isScinnaError(resp)) {
            setContext({...context, loaded: true, status: 'error', error: (resp as ScinnaError).Message})
            return;
        }

        setContext({...context, loaded: true, status: 'success', invites: resp as InviteCode[]})
    }

    const generate = async () => {
        const code = await apiCall<InviteCode>(token, {
            url: '/api/server/admin/invite',
            method: 'POST',
        });

        if (isScinnaError(code)) {
            // @TODO: do thing to error thing blabla
            return;
        }

        await refresh();
        setContext({...context, newlyGeneratedCode: (code as InviteCode).Code});
    };

    const remove = async (code: string) => {
        await apiCall(token, {
            url: '/api/server/admin/invite/' + code,
            method: 'DELETE',
        });

        await refresh();
    }

    useAsyncEffect(async () => {
        if (!context.loaded) {
            await refresh();
        }
    }, [context.loaded])

    return <InviteListContext.Provider value={{
        ...context,
        generate,
        refresh,
        remove
    }}>
        {children}
    </InviteListContext.Provider>
}

export function useInviteCode(): InviteListContextProps {
    return useContext<InviteListContextProps>(InviteListContext);
}