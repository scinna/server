import React from 'react';
import {BrowserHeader} from "../components/browser/BrowserHeader";
import {useBrowser} from "../context/BrowserProvider";
import {Icon} from "../components/browser/Icon";
import {useParams} from "react-router-dom";
import useAsyncEffect from "use-async-effect";

type RouteParams = {
    username: string;
    path?: string;
}

export function Browser() {
    const {username, path} = useParams<RouteParams>();
    const ctx = useBrowser();

    useAsyncEffect(async () => {
        console.log(username, path);
        await ctx.browse(username, path);
    }, [username, path])

    return <div>
        <BrowserHeader/>
        <div>
            {
                ctx.collection
                &&
                ctx.collection.Collections?.map(c => <Icon key={c.CollectionID} collection={c}/>)
            }
            {
                ctx.collection
                &&
                ctx.collection.Medias?.map(m => <Icon key={m.MediaID} media={m}/>)
            }
        </div>
    </div>;
}