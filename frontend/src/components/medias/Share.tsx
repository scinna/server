import {Media}           from "../../types/Media";
import React, {useState} from "react";

import classes     from "../../assets/scss/ShowPicture.module.scss";
import {Tab, Tabs} from "@material-ui/core";
import i18n        from "i18n-js";

type Props = {
    media: Media;
}

export const MediaShare = ({media}: Props) => {
    const [tab, setTab] = useState<number>(0);
    return <div className={classes.ShowPicture__Share}>
        <a className={classes.ShowPicture__Share__RawURL}
           href={window.location.protocol + '//' + window.location.hostname + '/' + media.MediaID}>Raw image URL</a>

        <Tabs value={tab}
              onChange={(_, val) => setTab(val)}
              indicatorColor="primary"
              variant="scrollable"
              scrollButtons="auto">
            <Tab label={"Markdown"}/>
            <Tab label={"Phpbb"}/>
        </Tabs>

        <div>
            {
                tab === 0
                &&
                <div>
                    Markdown
                </div>
            }
            {
                tab === 1
                &&
                <div>
                    Phpbb
                </div>
            }
        </div>
    </div>;
}