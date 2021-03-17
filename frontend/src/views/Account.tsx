import React, {useState} from 'react';
import {ProfileEditor} from "../components/account/ProfileEditor";
import {TokenLister} from "../components/account/TokenLister";
import {Tab, Tabs} from "@material-ui/core";
import i18n from "i18n-js";

import styles from '../assets/scss/Profile.module.scss';
import ShareX from "../components/account/ShareX";

function a11yProps(index: number) {
    return {
        id: `simple-tab-${index}`,
        'aria-controls': `simple-tabpanel-${index}`,
    };
}

export function Account() {
    const [currentTab, setCurrentTab] = useState<Number>(0);

    return <div className={styles.Account}>
        <Tabs value={currentTab}
              onChange={(_, val) => setCurrentTab(val)}
              indicatorColor="primary"
              centered>
            <Tab label={i18n.t('my_profile.account.tab_name')} {...a11yProps(0)}/>
            <Tab label={i18n.t('my_profile.tokens.tab_name')} {...a11yProps(1)}/>
            <Tab label={i18n.t('my_profile.sharex.tab_name')} {...a11yProps(2)}/>
        </Tabs>

        <div className={styles.Account__Tab}>
            {
                currentTab === 0
                &&
                <ProfileEditor/>
            }

            {
                currentTab === 1
                &&
                <TokenLister/>
            }

            {
                currentTab === 2
                &&
                <ShareX />
            }
        </div>

    </div>;
}