import React from 'react';

import {DropzoneArea} from 'material-ui-dropzone'

import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import Modal from '@material-ui/core/Modal';
import RadioGroup from '@material-ui/core/RadioGroup';
import Radio from '@material-ui/core/Radio';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import { LinearProgress, IconButton } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';

import UploadIcon from '@material-ui/icons/Publish';
import CloseIcon from '@material-ui/icons/Close';

import {APIUploadPicture, IUploadResponse} from '../api/Uploads';

import '../assets/UploadModal.scss'

interface IUploadProps {
    open: boolean,
    close: Function,
}

interface IState {
    Picture: {
        title: string,
        description: string,
        visibility: number,
        picture?: object|null
    }, 
    Buttons: {
        UploadEnabled: boolean
    },
    Upload: {
        InProgress: boolean, 
        Progress: number,
        UploadedOpen: boolean,
        UploadedResponse?: IUploadResponse|null,
    }
}

const initialState: IState = {
    Picture: {
        title: '',
        description: '',
        visibility: 0,
        picture: null,
    },
    Buttons: {
        UploadEnabled: false,
    },
    Upload: {
        InProgress: false,
        Progress: 0.0,
        UploadedOpen: false,
        UploadedResponse: null,
    }
}

export default function (props: IUploadProps) {
    const [state, setState] = React.useState(initialState);

    const closePopup = () => {
        if (!state.Upload.InProgress)
            props.close()
    }
    
    const cancelUpload: any = () => {

    }

    const handleUpload: any = (e: any) => {
        APIUploadPicture(state.Picture, (progress: number) => {
            setState({...state, Upload: { ...state.Upload, Progress: progress}})
        }, (response: IUploadResponse) => {
            setState({ ...state, Upload: {...state.Upload, UploadedOpen: true, UploadedResponse: response } })
        })
    }

    // @TODO: Type this everywhere there is a changeevent
    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const elt = e.currentTarget;
        const name = elt.getAttribute("name");

        if (name !== null)
            setState({ ...state, Picture: { ...state.Picture, [name]: elt.value }, Buttons: { UploadEnabled: state.Picture.picture != null && state.Picture.title.length > 0}})
    } 

    const handleFileChange = (files: object[]) => {
        const picture: object = files[0];
        setState({ ...state, Picture: {...state.Picture, picture}, Buttons: { ...state.Buttons, UploadEnabled: (picture != null && state.Picture.title.length > 0) } });
    }

    return (
        <Modal open={props.open} onClose={closePopup}>
            <div className="frame">
                <h3 className="title">Upload an image</h3>
                {/** @TODO: Handle the color correctly, with theme */}
                <Typography variant="body2" className="text">This will upload it to the current folder.</Typography>

                <form>
                    <div className="UploadZone">
                        <div className="infos">
                            <TextField name="title" label="Title" onChange={handleInputChange} value={state.Picture.title} required/>
                            <TextField name="description" label="Description" onChange={handleInputChange} value={state.Picture.description} multiline rowsMax="5"/>

                            <RadioGroup id="visibility" value={state.Picture.visibility.toString()} onChange={handleInputChange} row>
                                <FormControlLabel name="visibility" value="0" control={<Radio color="primary" />} label="Public" labelPlacement="bottom"/>
                                <FormControlLabel name="visibility" value="1" control={<Radio color="primary" />} label="Unlisted" labelPlacement="bottom"/>
                                <FormControlLabel name="visibility" value="2" control={<Radio color="primary" />} label="Private" labelPlacement="bottom"/>
                            </RadioGroup>
                        </div>
                        <div className="upload">
                            <DropzoneArea filesLimit={1} onChange={handleFileChange} acceptedFiles={[ "image/png", "image/jpeg", "image/gif", /** OK THANK YOU FIREFOX NOW FIX YOUR BUG PLEASE => */ ".jpeg", ".jpg" ]} />
                        </div>
                    </div>
                    <div className="ProgressZone">
                        <LinearProgress style={{height: '.75em'}} variant="determinate" value={state.Upload.Progress} className="progress" />
                        <IconButton className="upload" onClick={handleUpload} disabled={!state.Buttons.UploadEnabled || state.Upload.InProgress}>
                            <UploadIcon/>
                        </IconButton>
                        <IconButton className="cancel" onClick={cancelUpload}>
                            <CloseIcon />
                        </IconButton>
                    </div>
                </form>
                <ModalFileUploaded Open={state.Upload.UploadedOpen} OnClose={() => { setState({ ...state, Upload: { ...state.Upload, UploadedOpen: false } }) }} PictureData={state.Upload.UploadedResponse} />
            </div>
        </Modal>
    );
}

interface IUploadedData {
    Open: boolean,
    OnClose: any, // Ffs material-ui doing weird things
    PictureData?: IUploadResponse|null,
}

const stylesFileUploaded = makeStyles({
    title: {
        color: 'var(--above-fg-color)',
    }, 
    text: {
        color: 'var(--above-fg-color)',
    }
});

function ModalFileUploaded(props: IUploadedData) {
    const classes = stylesFileUploaded();

    return <Modal open={props.Open} onClose={props.OnClose}>
        <div className="frame">
            <h3 className={classes.title}>Picture uploaded!</h3>
                {/** @TODO: Handle the color correctly, with theme */}
                <Typography className={classes.text} variant="body2">Your picture has been uploaded.</Typography>
                <TextField value={window.location.protocol+"//"+window.location.hostname+"/"+props.PictureData?.URLID} /> {/* @TODO: Custom textfield with a clipboard option on the right */}
        </div>
    </Modal>
}