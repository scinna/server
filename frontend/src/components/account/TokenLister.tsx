import React, {useState} from 'react';
import {Token}                   from "../../types/Token";
import {Token as TokenComponent} from './Token';
import {useApiCall}              from "../../utils/useApi";
import {Loader}                  from "../Loader";
import {Button, DialogActions, DialogContent, DialogContentText, DialogTitle, Modal} from "@material-ui/core";
import { Dialog } from '@material-ui/core';
import i18n from "i18n-js";

export function TokenLister() {
    const tokens = useApiCall<Token[]>({ url: '/api/account/tokens' });
    const [revokedToken, setRevokedToken] = useState<string|null>(null);

    const revokeToken = () => {

        setRevokedToken(null);
    }

    return <div className="tokenLister">
        {
            tokens.status === 'pending'
            &&
                <Loader/>
        }
        {
            tokens.status === 'success'
            &&
                tokens.data.map(token => <TokenComponent key={token.Token} token={token} revokeToken={() => setRevokedToken(token.Token)}/>)
        }

        <Dialog open={revokedToken !== null} onClose={() => setRevokedToken(null)}>
            <DialogTitle>{i18n.t('my_profile.tokens.revoke_dialog.title')}</DialogTitle>
            <DialogContent>
                <DialogContentText>{i18n.t('my_profile.tokens.revoke_dialog.text')}</DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={() => setRevokedToken(null)} color="primary">
                    {i18n.t('my_profile.tokens.revoke_dialog.cancel')}
                </Button>
                <Button onClick={revokeToken} color="secondary">
                    {i18n.t('my_profile.tokens.revoke_dialog.revoke')}
                </Button>
            </DialogActions>
        </Dialog>
    </div>;
}