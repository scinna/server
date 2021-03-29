import {createContext, ReactNode, useContext, useState} from "react";

type Props = {
    children: ReactNode;
}

type ModalProps = {
    shown: boolean;
    node: ReactNode;
}

type ModalContextProps = ModalProps & {
    show: (dialog: ReactNode) => void;
    hide: () => void;
}

const initialState: ModalProps = {
    shown: false,
    node: null,
}

const ModalContext = createContext<ModalContextProps>({
    ...initialState,
    show: (node: ReactNode) => {},
    hide: () => {},
});

export default function ModalProvider({children}: Props) {
    const [context, setContext] = useState<ModalProps>(initialState);

    const show = (node: ReactNode) => setContext({...context, shown: true, node});
    const hide = () => setContext({...context, shown: false, node: null});

    return <ModalContext.Provider value={{
        ...initialState,
        show,
        hide,
    }}>
        {children}
        {context.node}
    </ModalContext.Provider>
}

export function useModal(): ModalContextProps {
    return useContext<ModalContextProps>(ModalContext);
}