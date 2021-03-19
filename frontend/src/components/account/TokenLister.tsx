import React, {useState} from 'react';
import {Token} from "../../types/Token";
import {Token as TokenComponent} from './Token';
import {apiCall, useApiCall} from "../../utils/useApi";
import {Loader} from "../Loader";
import {Button, DialogActions, DialogContent, DialogContentText, DialogTitle, Modal} from "@material-ui/core";
import {Dialog} from '@material-ui/core';
import i18n from "i18n-js";
import {useToken} from "../../context/TokenProvider";

import styles from '../../assets/scss/account/Profile.module.scss';

type TokenRevocation = {
    RevokedAt: string
}

export function TokenLister() {
    const {token} = useToken();
    const [isPending, setPending] = useState<boolean>(false);
    const [revokedToken, setRevokedToken] = useState<string | null>(null);

    // Double the request but meh, no clue on how to do it better
    const tokens = useApiCall<Token[]>({url: '/api/account/tokens'}, [isPending]);

    const revokeToken = async () => {
        setPending(true);

        await apiCall<TokenRevocation>(token, {
            url: '/api/account/tokens/' + revokedToken,
            method: 'DELETE',
        })

        setPending(false);
        setRevokedToken(null);
    }

    return <div className={styles.TabTokens}>
        {
            tokens.status === 'pending'
            &&
            <Loader/>
        }
        {
            tokens.status === 'success'
            &&
            tokens.data.map(token => <TokenComponent key={token.Token} token={token}
                                                     revokeToken={() => setRevokedToken(token.Token)}/>)
        }

        <Dialog open={revokedToken !== null} onClose={() => setRevokedToken(null)}>
            <DialogTitle>{i18n.t('my_profile.tokens.revoke_dialog.title')}</DialogTitle>
            <DialogContent>
                <DialogContentText>{i18n.t('my_profile.tokens.revoke_dialog.text')}</DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={() => setRevokedToken(null)} color="primary" disabled={isPending}>
                    {i18n.t('my_profile.tokens.revoke_dialog.cancel')}
                </Button>
                <Button onClick={revokeToken} color="secondary" disabled={isPending}>
                    {i18n.t('my_profile.tokens.revoke_dialog.revoke')}
                </Button>
            </DialogActions>
        </Dialog>
    </div>;
}