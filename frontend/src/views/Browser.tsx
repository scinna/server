import React from 'react';
import {BrowserHeader} from "../components/browser/BrowserHeader";
import {useBrowser} from "../context/BrowserProvider";
import {Icon} from "../components/browser/Icon";
import {useHistory, useParams} from "react-router-dom";
import useAsyncEffect from "use-async-effect";

import styles from '../assets/scss/browser/Browser.module.scss';
import {Loader} from "../components/Loader";
import {SpeedDial, SpeedDialAction, SpeedDialIcon} from "@material-ui/lab";
import {Description, Link} from "@material-ui/icons";
import i18n from "i18n-js";

type RouteParams = {
    username: string;
    path?: string;
}

export function Browser() {
    const history = useHistory();
    const {username, path} = useParams<RouteParams>();
    const [showSpeedDial, setShowSpeedDial] = React.useState<boolean>(false);
    const ctx = useBrowser();

    useAsyncEffect(async () => {
        await ctx.browse(username, path);
    }, [username, path])

    return <div className={styles.Browser + (ctx.pending ? ' ' + styles['Browser--Pending'] : '')}>
        <BrowserHeader/>
        {
            !ctx.pending
            &&
            <>
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

                <SpeedDial
                    className={styles.Browser__SpeedDial}
                    ariaLabel={'Speed dial'}
                    icon={<SpeedDialIcon/>}
                    onClose={() => setShowSpeedDial(false)}
                    onOpen={() => setShowSpeedDial(true)}
                    direction={'up'}
                    open={showSpeedDial}
                >

                    <SpeedDialAction
                        key={'upload-textbin'}
                        icon={<Description />}
                        tooltipTitle={i18n.t('dial.textbin')}
                    />

                    <SpeedDialAction
                        key={'url-shortener'}
                        icon={<Link />}
                        tooltipTitle={i18n.t('dial.url_shortener')}
                        onClick={() => history.push('/shortener')}
                    />

                </SpeedDial>
            </>
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