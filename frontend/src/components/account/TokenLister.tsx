import React, {useState}                                                              from 'react';
import {Token as TokenComponent}                                                      from './Token';
import {Loader}                                                                       from "../Loader";
import {Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle} from "@material-ui/core";
import i18n                                                                           from "i18n-js";

import styles             from '../../assets/scss/account/Profile.module.scss';
import {useAccountTokens} from "../../context/AccountTokenProvider";

export function TokenLister() {
    const [isPending, setPending] = useState<boolean>(false);
    const [revokedToken, setRevokedToken] = useState<string | null>(null);

    const {tokens, revoke, status, refresh} = useAccountTokens();

    const revokeToken = async () => {
        setPending(true);

        await revoke(revokedToken ?? "");
        await refresh();

        setPending(false);
        setRevokedToken(null);
    }

    return <div className={styles.TabTokens}>
        {
            status === 'pending'
            &&
            <Loader/>
        }
        {
            status === 'success'
            &&
            tokens?.map(token => <TokenComponent key={token.Token} token={token}
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