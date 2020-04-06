import React, { useReducer } from 'react';

import { useParams } from 'react-router';
import { Typography } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import { Skeleton } from '@material-ui/lab';

import { APIFetchMediaInfos, APIFetchPrivateMedia } from '../api/Medias';
import { AxiosMiddlishware } from '../api/Axios';

import MainReducer from '../reducers';
import { CtxInitialState } from '../context';

const initialState = {
    MediaLoaded: false,
    URLID: '',
    Title: '',
    Description: '',
    Visibility: 0,
    CreatedAt: new Date('2020-03-10T10:02:17.332307Z'),
    Creator: ''
};

const useStyles = makeStyles(theme => ({
    MediaContainer: {
        paddingTop: '2em',
        flex: 1,
        display: 'flex',
        flexDirection: 'column',
        flexWrap: 'wrap',
        height: '100%',
    },
    MediaWrapper: {
        background: 'var(--below-bg-color)',
        borderRadius: '.25em',
        width: '100%',
        flex: 1,
    },
    Media: {
        objectFit: 'contain',
        width: '100%',
        height: '100%',
    },
    MediaLoading: {
        flex: 1,
    },
    MediaInfo: {
        marginTop: '1.5em',
        background: 'var(--below-bg-color)',
        color: 'var(--below-fg-color)',
        borderRadius: '.25em',
        padding: '.75em',
        width: '100%',
    },
    MediaUploader: {
        color: theme.palette.primary.main,
        textDecoration: 'none',
    },
    MediaTitleBar: {
        display: 'flex',
        flexDirection: 'column',
        ['@media(min-width: 700px)']: {
            flexDirection: 'row',
        }
    },
    MediaTitle: {
        fontSize: '2em',
        maxWidth: '100%',
        ['@media (min-width: 700px)']: {
            flex: 1,
        }
    },
    MediaDescription: {
        fontSize: '1em',
        marginTop: '.75em',
    }
}));

export default function() {
    const classes = useStyles();
    const { pictID } = useParams();
    const [state, setState] = React.useState(initialState);
    const [global, dispatch] = useReducer(MainReducer, CtxInitialState);

    const login = AxiosMiddlishware(dispatch);
    if (login !== null) {
        return login;
    }

    React.useEffect(() => {
        APIFetchMediaInfos(pictID, (imageData: any) => {
            if (imageData.Visibility !== 2) {
                setState({ MediaLoaded: true, ...imageData, Creator: imageData.Creator.Username, CreatedAt: new Date(imageData.CreatedAt) });
                console.log("Public media")
            } else {
                APIFetchPrivateMedia(pictID, (image: string) => {
                    console.log("Private media")
                    setState({ MediaLoaded: true, ...imageData, URLID: image, Creator: imageData.Creator.Username, CreatedAt: new Date(imageData.CreatedAt) })
                });
            }
        });
    }, [])

    // Meh, I use image/jpeg even if the media is something else, we'll see later if that causes issues...
    return <div className={classes.MediaContainer}>
        <div className={classes.MediaWrapper}>
            { 
                state.MediaLoaded
                    ? <img className={classes.Media} src={state.Visibility !== 2 ? "/"+state.URLID : "data:image/jpeg;base64," + state.URLID} alt={state.Title} /> 
                    : <Skeleton className={`${classes.Media} ${classes.MediaLoading}`} variant="rect" height="100%" />
            }
        </div>
        <div className={classes.MediaInfo}>
            <div className={classes.MediaTitleBar}>
                { 
                    state.MediaLoaded 
                        ? <Typography className={classes.MediaTitle} variant="h1">{state.Title}</Typography> 
                        : <Skeleton className={classes.MediaTitle} variant="rect" width={400} height="1.5em" />
                }
                { 
                    state.MediaLoaded 
                        ? <span><a className={classes.MediaUploader} href="#">@{state.Creator}</a> > { state.CreatedAt.toLocaleDateString() }</span> 
                        : <Skeleton className={classes.MediaUploader} style={{ marginLeft: '20px', marginTop: '5px', maxWidth: 'calc(100% - 20px)', }} variant="rect" width={200} />}
            </div>
            { 
                state.MediaLoaded 
                    ? <Typography className={classes.MediaDescription} variant="body2">{state.Description}</Typography> 
                    : <Skeleton className={classes.MediaDescription} variant="rect" width="100%" height="2em" />
            }
        </div>
    </div>;
}