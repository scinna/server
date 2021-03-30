import React from 'react';

import styles             from '../../assets/scss/linkshortener/shortener.module.scss';
import {ShortenLink}      from "../../types/ShortenLink";
import {IconButton}       from "@material-ui/core";
import {Delete, FileCopy} from "@material-ui/icons";
import {useModal}         from "../../context/ModalProvider";
import {DeleteMedia}      from "../modals/DeleteMedia";
import {useShortenLink}   from "../../context/ShortenLinkProvider";
import i18n               from "i18n-js";

type Props = {
  link: ShortenLink;
};

export default function Link({link}: Props) {
    const modal = useModal();
    const {refresh} = useShortenLink();
    const fullLink = window.location.protocol + '//' + window.location.host + '/' + link.MediaID;

    const copyLink = () => {
        //@ts-ignore
        document.getElementById("LINK__" + link.MediaID).select();
        let copied = document.execCommand("copy");
    }

    return <div className={styles.Link}>
        <span className={styles.Link__Scinna}>{i18n.t('shortener.scinna_link')}: <a href={fullLink}>{fullLink}</a></span>
        <input className={styles.Link__HiddenTextfield} type="text" defaultValue={fullLink} readOnly={true} id={"LINK__" + link.MediaID}/>
        <a className={styles.Link__Real} href={link.Url}>{link.Url}</a>
        <span className={styles.Link__Views}>{i18n.t('shortener.amt_views')}: {link.AccessCount}</span>

        <div className={styles.Link__Icons}>
            <IconButton onClick={copyLink}>
                <FileCopy/>
            </IconButton>
            <IconButton onClick={() => modal.show(<DeleteMedia media={link} successCallback={refresh}/>)}>
                <Delete/>
            </IconButton>
        </div>
    </div>
}