import {IconButton, TextField} from "@material-ui/core";
import {
    ArrowBack as BackButton,
    ArrowForward as ForwardButton,
    Refresh,
    CreateNewFolder,
    CloudUpload as Upload
} from '@material-ui/icons';

import styles from '../../assets/scss/browser/Browser.module.scss';
import {useBrowser} from "../../context/BrowserProvider";
import {FolderCreator} from "./FolderCreator";
import {useState} from "react";
import {FileUploader} from "./FileUploader";

export const BrowserHeader = () => {
    const [folderCreationShown, setFolderCreationShown] = useState<boolean>(false);
    const [fileUploaderShown, setFileUploaderShown] = useState<boolean>(false);
    const { username, path, pending, refresh } = useBrowser();
    const fullPath = `/${username}/${path ? path : ''}`

    return <div className={styles.Browser__Header}>
        <IconButton disabled={pending} onClick={() => window.history.back()}>
            <BackButton/>
        </IconButton>
        <IconButton disabled={pending} onClick={() => window.history.forward()}>
            <ForwardButton/>
        </IconButton>
        <IconButton disabled={pending} onClick={() => refresh()}>
            <Refresh/>
        </IconButton>
        <TextField
            className={styles.Browser__Header__AddressBar}
            value={fullPath}
            disabled={true}
        />
        {/**
            Temporary, just so we can't create nested collections
         **/}
        <IconButton onClick={() => setFolderCreationShown(true)} disabled={pending || (path !== undefined && path?.length > 0)}>
            <CreateNewFolder/>
        </IconButton>
        <IconButton onClick={() => setFileUploaderShown(true)} disabled={pending}>
            <Upload/>
        </IconButton>

        <FolderCreator shown={folderCreationShown} onClose={() => setFolderCreationShown(false)}/>
        <FileUploader shown={fileUploaderShown} onClose={() => setFileUploaderShown(false)}/>
    </div>;
}