import React from 'react';

import { useLocation } from 'react-router-dom';

import { makeStyles, Typography, Fab } from '@material-ui/core';
import AddPhotoAlternateIcon from '@material-ui/icons/AddPhotoAlternate';
import NewFolderIcon from '@material-ui/icons/AddToPhotos';

import { useStateValue } from '../context';
import NotFound from '../components/NotFound';
import Login from '../components/Login';
import UploadComponent from '../components/UploadComponent';
import FileBrowser from '../components/FileBrowser';


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

const initialState = {
    UploadModal: false,
    FolderModal: false,
};

export default function() {
    const classes = useStyles();
    const location = useLocation();
    const [modal, setModal] = React.useState(initialState);


    console.log(location);

    //@ts-ignore
    const [global] = useStateValue();
    const userLoggedIn = global.User.Username.length > 0 && global.User.Token.length > 0;

    const closeModal = (currModal: string) => () => {
        setModal({
            ...modal,
            [currModal]: false,
        });
    };

    let page;
    if (userLoggedIn) {
        page = <div>
            <Typography className={classes.title} variant="h4" component="h1">My content</Typography>
            <FileBrowser />

            <div className={classes.fab}>
                <Fab color="primary" aria-label="Upload picture" onClick={() => setModal({ ...modal, UploadModal: true })}>
                    <AddPhotoAlternateIcon />
                </Fab>
                <Fab className={classes.fabIcons} color="primary" aria-label="Create folder" onClick={() => setModal({ ...modal, FolderModal: true })}>
                    <NewFolderIcon />
                </Fab>
            </div>

            <UploadComponent open={modal.UploadModal} close={closeModal("UploadModal")} />
        </div>
    } else {
        if (location.pathname !== '/')
            page = <NotFound />
        else
            page = <Login />
    }

    return page;
}