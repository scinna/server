import React from 'react';

import styles from '../../assets/scss/linkshortener/shortener.module.scss';
import {ShortenLink} from "../../types/ShortenLink";
import {IconButton} from "@material-ui/core";
import {Delete, FileCopy} from "@material-ui/icons";

type Props = {
  link: ShortenLink;
};

export default function Link({link}: Props) {
    const fullLink = window.location.protocol + '//' + window.location.host + '/' + link.MediaID;

    return <div className={styles.Link}>
        <span className={styles.Link__Scinna}>Lien Scinna: <a href={fullLink}>{fullLink}</a></span>
        <a className={styles.Link__Real} href={link.Url}>{link.Url}</a>
        <span className={styles.Link__Views}>Views: {link.AccessCount}</span>

        <div className={styles.Link__Icons}>
            <IconButton>
                <FileCopy/>
            </IconButton>
            <IconButton>
                <Delete/>
            </IconButton>
        </div>
    </div>
}