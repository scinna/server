import React from 'react';

import {makeStyles} from '@material-ui/core/styles';

import File from '../model/File';

import FolderIcon from '../assets/folder_icon.png';
import BackIcon from '../assets/back_icon.png';

const useStyles = makeStyles({
    fullIcon: {
        color: '#fff',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        width: '128px',
        cursor: 'pointer'
    },
    icon: {
        width: '3rem',
        height: '3rem',
        objectFit: 'cover',
    },
    text: {
        marginTop: '.5rem',
        textAlign: 'center',
    }
})

interface FileProps {
    className?: string,
    isGoBack?: boolean,
    onClick?: any,
    file?: File
}

export default function(props: FileProps) {
    const classes = useStyles();

    let icon;
    let text;

    if (props.isGoBack) {
        icon = BackIcon;
        text = "Back";
    } else if (props.file && props.file.isFolder) {
        icon = FolderIcon;
    } else {
        // First result typing "image" on google. Just to poke around
        icon = "https://media.gettyimages.com/photos/colorful-powder-explosion-in-all-directions-in-a-nice-composition-picture-id890147976?s=612x612"
    }

    return <div className={`${classes.fullIcon} ${props.className}`} onClick={props.onClick}>
        <img className={classes.icon} alt="File icon" src={icon} />
        <span className={classes.text}>{props.file ? props.file.filename : text}</span>
    </div>
}