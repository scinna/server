import React                from 'react';
import i18n                 from 'i18n-js';
import {Token as TokenType} from '../../types/Token';
import {displayDate}        from "../../utils/DateUtils";

import styles from '../../assets/scss/_Token.module.scss';

type TokenProps = {
    token: TokenType;
}

export function Token({ token }: TokenProps) {
    return <div className={styles.Token + " " + (token.RevokedAt !== null ? styles.Token__Revoked : "")}>
        <p>{i18n.t('my_profile.tokens.logged_at')} {displayDate(token.CreatedAt)}</p>
        <p>{i18n.t('my_profile.tokens.last_seen')}: {displayDate(token.LastSeen)}</p>
        {
            token.RevokedAt
                && <p>{i18n.t('my_profile.tokens.revoked_at')} {displayDate(token.RevokedAt)}</p>
        }
        <p>{token.LoginIP}</p>
        <a href="#">Logout</a>
    </div>;
}