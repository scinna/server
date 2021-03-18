import React, {useState} from 'react';
import {ProfileEditor} from "../components/account/ProfileEditor";
import {TokenLister} from "../components/account/TokenLister";
import {Tab, Tabs} from "@material-ui/core";
import i18n from "i18n-js";

import styles from '../assets/scss/account/Profile.module.scss';
import ShareX from "../components/account/ShareX";

export function Account() {
    const [currentTab, setCurrentTab] = useState<Number>(0);

    return <div className={styles.Tabbed}>
        <Tabs value={currentTab}
              onChange={(_, val) => setCurrentTab(val)}
              indicatorColor="primary"
              variant="scrollable"
              scrollButtons="auto">
            <Tab label={i18n.t('my_profile.account.tab_name')}/>
            <Tab label={i18n.t('my_profile.tokens.tab_name')}/>
            <Tab label={i18n.t('my_profile.sharex.tab_name')}/>
        </Tabs>

        <div className={styles.Tabbed__Tab}>
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