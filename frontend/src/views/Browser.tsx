import React from 'react';
import {BrowserHeader} from "../components/browser/BrowserHeader";
import {useBrowser} from "../context/BrowserProvider";
import {Icon} from "../components/browser/Icon";
import {useParams} from "react-router-dom";
import useAsyncEffect from "use-async-effect";

import styles from '../assets/scss/browser/Browser.module.scss';
import {Loader} from "../components/Loader";

type RouteParams = {
    username: string;
    path?: string;
}

export function Browser() {
    const {username, path} = useParams<RouteParams>();
    const ctx = useBrowser();

    useAsyncEffect(async () => {
        await ctx.browse(username, path);
    }, [username, path])

    return <div className={styles.Browser + (ctx.pending ? ' ' + styles['Browser--Pending'] : '')}>
        <BrowserHeader/>
        {
            !ctx.pending
            &&
            <div className={styles.Browser__IconList}>
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

        }

        {
            ctx.pending
            &&
            <div className={styles.Browser__Pending}>
                <Loader/>
            </div>
        }
    </div>;
}