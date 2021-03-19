import styles                                                                         from '../../assets/scss/server/ServerSettings.module.scss';
import {apiCall, useApiCall}                                                          from "../../utils/useApi";
import {Loader}                                                                       from "../Loader";
import {InviteCode}                                                                   from "./InviteCode";
import {InviteCodeGenerator}                                                          from "./InviteCodeGenerator";
import React, {useState}                                                              from "react";
import {Button, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle} from "@material-ui/core";
import i18n                                                                           from "i18n-js";
import {useToken}                                                                     from "../../context/TokenProvider";
import {isScinnaError}                                                                from "../../types/Error";
import {useInviteCode}                                                                from "../../context/InviteCodeProvider";

export function TabInviteCodes() {
    const {token} = useToken();
    const {invites, status} = useInviteCode();

    const [isPending, setPending] = useState<boolean>(false);
    const [toDeleteInvite, setToDeleteInvite] = useState<InviteCode | null>(null);

    const deleteInvite = async (invite: InviteCode|null) => {
        if (invite === null) {
            return
        }

        setPending(true);

        const resp = await apiCall(token, {
            url: '/api/server/admin/invite/' + invite.Code,
            method: 'DELETE',
        })

        if (isScinnaError(resp)) {
            // @TODO: Show a message
            console.log(resp);
            setPending(false);
            setToDeleteInvite(null);
            return;
        }

        setPending(false);
        setToDeleteInvite(null);
    }

    return <div className={styles.InviteCodesTab}>
        <InviteCodeGenerator/>

        <div>
            {
                status === 'pending'
                &&
                <Loader/>
            }
            {
                status === 'success'
                &&
                invites?.map(invite => <InviteCode invite={invite} askForDeletion={() => setToDeleteInvite(invite)}/>)
            }
            {
                status === 'error'
                &&
                <p>invites.error.Message</p>
            }

            <Dialog open={toDeleteInvite !== null} onClose={() => setToDeleteInvite(null)}>
                <DialogTitle>{i18n.t('server_settings.invite.delete_dialog.title')}</DialogTitle>
                <DialogContent>
                    <DialogContentText>{i18n.t('server_settings.invite.delete_dialog.text')}</DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={() => setToDeleteInvite(null)} color="primary" disabled={isPending}>
                        {i18n.t('server_settings.invite.delete_dialog.cancel')}
                    </Button>
                    <Button onClick={() => deleteInvite(toDeleteInvite)} color="secondary" disabled={isPending}>
                        {i18n.t('server_settings.invite.delete_dialog.remove')}
                    </Button>
                </DialogActions>
            </Dialog>
        </div>
    </div>;
}