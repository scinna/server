import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    Typography
}                                   from "@material-ui/core";
import React, {useState}            from "react";
import {useModal}                   from "../../context/ModalProvider";
import {apiCall}                    from "../../utils/useApi";
import {useToken}                   from "../../context/TokenProvider";
import {isScinnaError, ScinnaError} from "../../types/Error";
import {useBrowser}                 from "../../context/BrowserProvider";
import i18n                         from "i18n-js";
import {Media}                      from "../../types/Media";
import {ShortenLink}                from "../../types/ShortenLink";

type Props = {
    media: Media | ShortenLink;
    successCallback?: () => void;
}

export function DeleteMedia({media, successCallback}: Props) {
    const {token} = useToken();
    const {refresh} = useBrowser();
    const [pending, setPending] = useState<boolean>(false);
    const [error, setError] = useState<string>('');
    const {hide} = useModal();

    const deleteMedia = async () => {
        await setPending(true);
        await setError('');

        const response = await apiCall(token, {
            url: '/' + media.MediaID,
            method: 'DELETE',
        });

        if (isScinnaError(response)) {
            let err = response as ScinnaError;
            if (err.status !== 410) {
                await setError((response as ScinnaError).Message);
                await setPending(false);

                return;
            }
        }

        await refresh();
        await hide();

        if (successCallback) {
            successCallback();
        }
    };

    const isShortenLink = !media.hasOwnProperty('Title');

    return <Dialog open={true} onClose={hide}>
        {
            !isShortenLink
            &&
            <DialogTitle>{i18n.t('browser.modals.remove_media.title')} {(media as Media).Title}</DialogTitle>
        }
        {
            isShortenLink
            &&
            <DialogTitle>{i18n.t('browser.modals.remove_link.title')}</DialogTitle>
        }
        <DialogContent>
            {
                !isShortenLink
                &&
                <DialogContentText>{i18n.t('browser.modals.remove_media.text')}</DialogContentText>
            }
            {
                isShortenLink
                &&
                <DialogContentText>{i18n.t('browser.modals.remove_link.text')} {(media as ShortenLink).Url}</DialogContentText>
            }
            {
                error.length > 0
                &&
                <Typography color="secondary">{error}</Typography>
            }
        </DialogContent>
        <DialogActions>
            <Button onClick={hide} color="primary" disabled={pending}>
                {i18n.t('browser.modals.remove_media.cancel')}
            </Button>
            <Button onClick={deleteMedia} color="secondary" disabled={pending}>
                {i18n.t('browser.modals.remove_media.delete')}
            </Button>
        </DialogActions>
    </Dialog>
}