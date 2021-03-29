import React, {createContext, ReactNode, useContext, useState} from "react";
import {Collection}                                            from "../types/Collection";
import {Media}                                                 from "../types/Media";

type Props = {
    children: ReactNode;
}

type IconContextProps = {
    mouseX: number | null;
    mouseY: number | null;
    collection: Collection | null;
    media: Media | null;
}

const initialState = {
    mouseX: null,
    mouseY: null,
    collection: null,
    media: null,
}

type IconContextMenuContextProps = IconContextProps & {
    show: (collection: Collection|null, media: Media|null) => (mouseEvent: React.MouseEvent) => void;
    hide: () => void;
    isVisible: () => boolean;
}

const IconContextMenuContext = createContext<IconContextMenuContextProps>({
    ...initialState,
    show: (c, m) => (e) => {
    },
    hide: () => {
    },
    isVisible: () => false,
})

export default function IconContextMenuProvider({children}: Props) {
    const [context, setContext] = useState<IconContextProps>(initialState);

    const show = (collection: Collection|null, media: Media|null = null) => (mouseEvent: React.MouseEvent) => {
        mouseEvent.preventDefault();
        setContext({...context, mouseX: mouseEvent.clientX - 2, mouseY: mouseEvent.clientY - 4, collection, media})
    }

    const hide = () => {
        setContext(initialState);
    }

    const isVisible = () => context.mouseX !== null && context.mouseY !== null && (context.collection !== null || context.media !== null);

    return <IconContextMenuContext.Provider value={{
        ...context,
        show,
        hide,
        isVisible,
    }}>
        {children}
    </IconContextMenuContext.Provider>
}

export function useIconContextMenu(): IconContextMenuContextProps{
    return useContext<IconContextMenuContextProps>(IconContextMenuContext);
}