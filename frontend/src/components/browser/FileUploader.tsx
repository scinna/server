import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    InputLabel, LinearProgress,
    TextareaAutosize,
    TextField
} from '@material-ui/core';
import i18n from 'i18n-js';
import React, {useState} from 'react';
import {useToken} from '../../context/TokenProvider';
import {Controller, useForm} from 'react-hook-form';
import {VisibilityDropDown} from '../VisibilityDropDown';
import {Dropzone} from '../Dropzone';

import styles              from '../../assets/scss/browser/Browser.module.scss';
import {CopiableTextfield} from "../CopiableTextfield";
import {useServerConfig}   from "../../context/ServerConfigProvider";
import {useBrowser}        from "../../context/BrowserProvider";

type Props = {
    shown: boolean;
    onClose: () => void;
}

type IFormInputs = {
    title: string;
    description: string;
    visibility: number;
}

export function FileUploader({shown, onClose}: Props) {
    const {token} = useToken();
    const {Config} = useServerConfig();
    const {refresh} = useBrowser();
    const [status, setStatus] = useState<string>('');
    const [progress, setProgress] = useState<number>(0);
    const {control, handleSubmit, reset} = useForm<IFormInputs>();

    const [selectedFile, setSelectedFile] = useState<File>();
    const [uploadedId, setUploadedId] = useState<string>('');

    const resetAndClose = () => {
        onClose();
        reset();
        setSelectedFile(undefined);
        setUploadedId('');
        setStatus('');
        setProgress(0);
    }

    // Ugly but it appears that fetch can not handle progression for now
    const onSubmit = async (data: IFormInputs) => {
        if (!selectedFile) {
            return;
        }
        await setProgress(0);
        await setStatus('pending');

        let fd = new FormData();
        Object.entries(data).forEach(([key, val]) => {
            fd.append(key, '' + val);
        });
        fd.append('picture', selectedFile);

        let request = new XMLHttpRequest();
        request.open('POST', '/api/upload');
        request.setRequestHeader('Authorization', 'Bearer ' + token);

        request.upload.addEventListener('progress', (e) => {
            setProgress((e.loaded / e.total) * 100);
        });

        request.addEventListener('load', (e) => {
            if (request.status !== 201) {
                try {
                    const resp = JSON.parse(request.response);
                    setStatus(resp.Message);
                } catch (e) {
                    setStatus(i18n.t('errors.unknown'));
                }

                return;
            }

            const resp = JSON.parse(request.response);
            const id = resp.MediaID;
            setUploadedId(id);

            reset();
            refresh();
            setSelectedFile(undefined);
            setStatus('success');
        });

        request.send(fd);
    }

    return <Dialog open={shown} onClose={resetAndClose}>
        {
            status !== 'success'
            &&
            <form onSubmit={handleSubmit(onSubmit)}>
                <DialogTitle>{i18n.t('browser.file_uploader.screen_1.title')}</DialogTitle>
                <DialogContent>
                    <InputLabel>{i18n.t('browser.file_uploader.screen_1.file_name')}: </InputLabel>
                    <Controller
                        name={"title"}
                        control={control}
                        defaultValue=""
                        render={({onChange, value}) =>
                            <TextField
                                onChange={onChange}
                                value={value}
                                disabled={status === 'pending'}
                                fullWidth={true}
                                required
                            />}
                    />
                    <InputLabel className={styles.FileUploader__VisibilityLabel}>
                        {i18n.t('browser.file_uploader.screen_1.description')}:
                    </InputLabel>
                    <Controller
                        name={"description"}
                        control={control}
                        defaultValue={''}
                        render={({onChange, value}) => <TextareaAutosize
                            className={styles.FileUploader__Description}
                            disabled={status === 'pending'}
                            rowsMin={3}
                            value={value}
                            onChange={onChange}
                        />
                        }
                    />
                    <InputLabel className={styles.FileUploader__VisibilityLabel}>
                        {i18n.t('browser.file_uploader.screen_1.visibility')}:
                    </InputLabel>
                    <Controller
                        name={"visibility"}
                        control={control}
                        defaultValue={0}
                        render={({onChange, value}) => <VisibilityDropDown
                            selectedVisibility={value}
                            setSelectedVisibility={onChange}
                            disabled={status === 'pending'}
                        />
                        }
                    />
                    <Dropzone onFileSelected={f => setSelectedFile(f)}/>
                    <LinearProgress className={styles.FileUploader__Progress} variant="determinate" value={progress}/>

                    {
                        status.length > 0 && status !== 'pending'
                        &&
                        <span className={styles.FileUploader__Error}>{status}</span>
                    }

                </DialogContent>
                <DialogActions>
                    <Button onClick={resetAndClose}
                            disabled={status === 'pending'}>{i18n.t('browser.file_uploader.screen_1.cancel')}</Button>
                    <Button color="primary" type="submit"
                            disabled={status === 'pending'}>{i18n.t('browser.file_uploader.screen_1.upload')}</Button>
                </DialogActions>
            </form>
        }
        {
            status === 'success'
            &&
                <>
                    <DialogTitle>{i18n.t('browser.file_uploader.screen_2.title')}</DialogTitle>
                    <DialogContent>
                        <p>{i18n.t('browser.file_uploader.screen_2.text')}</p>
                        <p className={styles.FileUploader__Success__Link}>{i18n.t('browser.file_uploader.screen_2.scinna_link')}</p>
                        <CopiableTextfield value={Config.WebURL+"app/"+uploadedId}/>
                        <p>{i18n.t('browser.file_uploader.screen_2.raw_link')}</p>
                        <CopiableTextfield value={Config.WebURL+uploadedId}/>
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={resetAndClose}>{i18n.t('browser.file_uploader.screen_2.close')}</Button>
                    </DialogActions>
                </>
        }
    </Dialog>
}