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
import {Collection}                 from "../../types/Collection";
import {apiCall}                    from "../../utils/useApi";
import {useToken}                   from "../../context/TokenProvider";
import {isScinnaError, ScinnaError} from "../../types/Error";
import {useBrowser}                 from "../../context/BrowserProvider";
import i18n                         from "i18n-js";

type Props = {
    collection: Collection;
    successCallback: () => void;
}

export function DeleteCollection({collection, successCallback}: Props) {
    const {token} = useToken();
    const {getCurrentPath, refresh} = useBrowser();
    const [pending, setPending] = useState<boolean>(false);
    const [error, setError] = useState<string>('');
    const {hide} = useModal();

    const deleteCollection = async () => {
        await setPending(true);
        await setError('');

        const response = await apiCall(token, {
            url: '/api/browse/' + getCurrentPath() + collection.Title,
            method: 'DELETE',
        });

        if (isScinnaError(response)) {
            let err = response as ScinnaError;
            if (err.status != 410) {
                await setError((response as ScinnaError).Message);
                await setPending(false);

                return;
            }
        }

        await refresh();
        await hide();

        successCallback();
    };

    return <Dialog open={true} onClose={hide}>
        <DialogTitle>{i18n.t('browser.modals.remove_collection.title')} {collection.Title}</DialogTitle>
        <DialogContent>
            <DialogContentText>{i18n.t('browser.modals.remove_collection.text')}</DialogContentText>
            {
                error.length > 0
                &&
                <Typography color="secondary">{error}</Typography>
            }
        </DialogContent>
        <DialogActions>
            <Button onClick={hide} color="primary" disabled={pending}>
                {i18n.t('browser.modals.remove_collection.cancel')}
            </Button>
            <Button onClick={deleteCollection} color="secondary" disabled={pending}>
                {i18n.t('browser.modals.remove_collection.delete')}
            </Button>
        </DialogActions>
    </Dialog>
}