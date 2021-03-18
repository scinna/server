import React from 'react';
import styles from "../../assets/scss/server/ServerSettings.module.scss";
import {IconButton} from "@material-ui/core";
import CloseIcon from "@material-ui/icons/Close";

export type InviteCode = {
    Code: string;
    Author: string;
    GeneratedAt: Date;
    Used: boolean;
};

type Props = {
    invite: InviteCode,
    askForDeletion: () => void
}

export function InviteCode({invite, askForDeletion}: Props) {
    return <div className={
        styles.InviteCodesTab__List__Item + " " +
        (invite.Used ? styles.InviteCodesTab__List__Item__Used : "")
    }>
        <div className={styles.InviteCodesTab__List__Item__Infos}>
            <p>{invite.Code}</p>
            <p>Generated by {invite.Author}</p>
            <p>On {invite.GeneratedAt}</p>
        </div>
        <IconButton color="primary" aria-label="logout picture" component="span" disabled={invite.Used} onClick={askForDeletion}>
            <CloseIcon/>
        </IconButton>
    </div>
}