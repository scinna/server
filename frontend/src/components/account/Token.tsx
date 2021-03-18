import React, {useState} from 'react';
import i18n from 'i18n-js';
import {Token as TokenType} from '../../types/Token';
import {displayDate} from "../../utils/DateUtils";
import {IconButton} from "@material-ui/core";
import ExitToApp from '@material-ui/icons/ExitToApp';

import styles from '../../assets/scss/account/_Token.module.scss';

type TokenProps = {
    token: TokenType;
    revokeToken: () => void;
}

export function Token({token, revokeToken}: TokenProps) {
    return <div className={styles.Token + " " + (token.RevokedAt !== null ? styles.Token__Revoked : "")}>
        <div className={styles.Token__Infos}>
            <p>{i18n.t('my_profile.tokens.logged_at')} {displayDate(token.CreatedAt)}</p>
            <p>{i18n.t('my_profile.tokens.last_seen')}: {displayDate(token.LastSeen)}</p>
            {
                token.RevokedAt
                && <p>{i18n.t('my_profile.tokens.revoked_at')} {displayDate(token.RevokedAt)}</p>
            }
            <p>{token.LoginIP}</p>
        </div>
        <IconButton color="primary" aria-label="logout picture" component="span" disabled={token.RevokedAt !== null} onClick={revokeToken}>
            <ExitToApp/>
        </IconButton>
    </div>;
}