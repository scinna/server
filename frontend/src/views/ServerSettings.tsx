import React, {useState}  from 'react';
import {Tab, Tabs}        from "@material-ui/core";
import i18n               from 'i18n-js';
import styles             from '../assets/scss/server/ServerSettings.module.scss';
import {TabInviteCodes}   from "../components/server/TabInviteCodes";
import InviteCodeProvider from "../context/InviteCodeProvider";

export function ServerSettings() {
    const [currentTab, setCurrentTab] = useState<Number>(0);
    return <div className={styles.Tabbed}>
        <Tabs value={currentTab}
              onChange={(_, val) => setCurrentTab(val)}
              indicatorColor="primary"
              variant="scrollable"
              scrollButtons="auto">
            <Tab label={i18n.t('server_settings.invite.tab_name')}/>
        </Tabs>

        <div className={styles.Tabbed__Tab}>
            {
                currentTab === 0
                &&
                <InviteCodeProvider>
                    <TabInviteCodes/>
                </InviteCodeProvider>
            }
        </div>
    </div>;
}