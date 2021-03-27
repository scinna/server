import {Media} from "../../types/Media";
import React, {useState} from "react";

import classes from "../../assets/scss/ShowPicture.module.scss";
import {Tab, Tabs} from "@material-ui/core";

type Props = {
    media: Media;
}

export const MediaShare = ({media}: Props) => {
    const [tab, setTab] = useState<number>(0);

    const rootUrl = window.location.protocol + '//' + window.location.hostname + '/';
    const rawUrl = rootUrl + media.MediaID;
    const appUrl = rootUrl + 'app/' + media.MediaID;

    return <div className={classes.ShowPicture__Share}>
        <Tabs value={tab}
              onChange={(_, val) => setTab(val)}
              indicatorColor="primary"
              variant="scrollable"
              scrollButtons="auto">
            <Tab label={"Raw URL"}/>
            <Tab label={"Markdown"}/>
            <Tab label={"HTML"}/>
            <Tab label={"Phpbb"}/>
        </Tabs>

        <div className={classes.ShowPicture__Share__Content}>
            {
                tab === 0
                &&
                <a className={classes.ShowPicture__Share__RawURL}
                   href={rawUrl}>{rawUrl}</a>
            }
            {
                tab === 1
                &&
                <pre>
                    <code>
                        [![{media.Title}]({rawUrl})]({appUrl})
                    </code>
                </pre>
            }
            {
                tab === 2
                &&
                <pre>
                    &lt;a href="{rawUrl}"&gt;
                        &lt;img src="{rawUrl}" alt="{media.Title}" /&gt;
                    &lt;/a&gt;
                </pre>
            }
            {
                tab === 3
                &&
                <pre>
                    [url={appUrl}][img]{rawUrl}[/img][/url]
                </pre>
            }
        </div>
    </div>;
}