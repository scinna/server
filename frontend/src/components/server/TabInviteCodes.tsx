import {Loader}                                                                       from "../Loader";
import {InviteCode, InviteCodeComponent}                                              from "./InviteCode";
import React, {useState}                                                              from "react";
import {Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle} from "@material-ui/core";
import i18n                                                                           from "i18n-js";
import {useInviteCode}                                                                from "../../context/InviteCodeProvider";

import styles from '../../assets/scss/server/ServerSettings.module.scss';

export function TabInviteCodes() {
    const {invites, status, error, remove, generate, refresh} = useInviteCode();
    const [toDeleteInvite, setToDeleteInvite] = useState<InviteCode | null>(null);

    return <div className={styles.InviteCodesTab}>
        <Button className={styles.InviteCodesTab__GenerateButton}
                onClick={async () => {
                    await generate();
                    await refresh();
                }}>
            {i18n.t('server_settings.invite.generate')}
        </Button>

        <div>
            {
                status === 'pending'
                &&
                <Loader/>
            }
            {
                status === 'success'
                &&
                invites?.map(invite => <InviteCodeComponent invite={invite}
                                                            askForDeletion={() => setToDeleteInvite(invite)}/>)
            }
            {
                status === 'error'
                &&
                <p>{error}</p>
            }

            <Dialog open={toDeleteInvite !== null} onClose={() => setToDeleteInvite(null)}>
                <DialogTitle>{i18n.t('server_settings.invite.delete_dialog.title')}</DialogTitle>
                <DialogContent>
                    <DialogContentText>{i18n.t('server_settings.invite.delete_dialog.text')}.</DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={() => setToDeleteInvite(null)} color="primary" disabled={status === 'pending'}>
                        {i18n.t('server_settings.invite.delete_dialog.cancel')}
                    </Button>
                    <Button onClick={async () => {
                        await remove(toDeleteInvite?.Code ?? "");
                        setToDeleteInvite(null);
                    }} color="secondary"
                            disabled={status === 'pending'}>
                        {i18n.t('server_settings.invite.delete_dialog.remove')}
                    </Button>
                </DialogActions>
            </Dialog>
        </div>
    </div>;
}