import React from 'react';
import {Button, TextField} from "@material-ui/core";
import i18n                from "i18n-js";
import {useInviteCode} from "../../context/InviteCodeProvider";

import styles          from '../../assets/scss/server/_InviteGenerator.module.scss';

export function InviteCodeGenerator() {
    const { generate, newlyGeneratedCode } = useInviteCode();

    return <div className={styles.InviteGenerator}>
        <TextField disabled={true} label={i18n.t('server_settings.invite.code')} value={newlyGeneratedCode}/>
        <Button onClick={generate}>{i18n.t('server_settings.invite.generate')}</Button>
    </div>;
}