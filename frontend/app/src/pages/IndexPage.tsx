import React from 'react';
import { useStateValue } from '../context';

import Login from '../components/Login';
import FileBrowser from '../components/FileBrowser';

import { makeStyles, Typography, Fab } from '@material-ui/core';
import AddPhotoAlternateIcon from '@material-ui/icons/AddPhotoAlternate';
import NewFolderIcon from '@material-ui/icons/AddToPhotos';

const useStyles = makeStyles(theme => ({
    title: {
        color: '#cececf', // @Todo find Material-ui's theme property for titles and put it there
        marginLeft: '.5em',
        marginBottom: '.25em',
    },
    fab: {
        position: 'fixed',
        right: '2em',
        bottom: '2em',
    },
    fabIcons: {
        marginLeft: '1em',
    }
}));

export default function() {
    const classes = useStyles();

    //@ts-ignore
    const [global] = useStateValue();

    const userLoggedIn = global.User.Username.length > 0 && global.User.Token.length > 0;

    let page;

    if (userLoggedIn) {
        page = <div>
            <Typography className={classes.title} variant="h4" component="h1">My content</Typography>
            <FileBrowser />

            <div className={classes.fab}>
                <Fab color="primary" aria-label="add">
                    <AddPhotoAlternateIcon />
                </Fab>
                <Fab className={classes.fabIcons} color="primary" aria-label="add">
                    <NewFolderIcon />
                </Fab>
            </div>
        </div>
    } else {
        page = <Login />
    }

    return page;
}