import React from 'react';
import i18n from 'i18n-js';
import {Token as TokenType}  from '../types/Token';

type TokenProps = {
    token: TokenType;
}

export function Token({ token }: TokenProps) {
    return <div className={"token " + (token.RevokedAt !== null ? "revoked" : "")}>
        <p>{i18n.t('my_profile.loggedAt')} {token.CreatedAt}</p>
        <p>{i18n.t('my_profile.last_seen')}: {token.LastSeen ?? i18n.t('my_profile.never')}</p>
        <p>{token.LoginIP}</p>
        <a href="#">Logout</a>
    </div>;
}