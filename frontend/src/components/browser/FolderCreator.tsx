import {
    Button,
    Dialog,
    DialogActions,
    DialogContent, DialogContentText,
    DialogTitle, InputLabel,
    TextField
} from "@material-ui/core";
import i18n from "i18n-js";
import React, {useState} from "react";
import {useForm, Controller} from "react-hook-form";
import {apiCall} from "../../utils/useApi";
import {useToken} from "../../context/TokenProvider";
import {useBrowser} from "../../context/BrowserProvider";
import {isScinnaError, ScinnaError} from "../../types/Error";

type Props = {
    shown: boolean;
    onClose: () => void;
}

type IFormInputs = {
    name: string;
}

export function FolderCreator({shown, onClose}: Props) {
    const {username, path, refresh} = useBrowser();
    const {token} = useToken();
    const [pending, setPending] = useState<boolean>(false);
    const [error, setError] = useState<String>("");
    const {control, handleSubmit} = useForm<IFormInputs>();

    const onSubmit = async (data: IFormInputs) => {
        await setPending(true);

        const response = await apiCall(token, {
            url: '/api/browse/' + username + '/' + path + (path && path?.length > 0 ? '/' : '') + data.name,
            method: 'POST',
            data: {
                Visibility: 0,
            }
        });

        if (isScinnaError(response)) {
            await setError((response as ScinnaError).Message);
            await setPending(false);
            return;
        }

        await setPending(false);
        await setError('');
        await refresh();
        await onClose();
    };

    return <Dialog open={shown} onClose={onClose}>
        <form onSubmit={handleSubmit(onSubmit)}>
            <DialogTitle>{i18n.t('browser.create_folder.title')}</DialogTitle>
            <DialogContent>
                <InputLabel htmlFor="name">{i18n.t('browser.create_folder.folder_name')}</InputLabel>
                <Controller
                    name={"name"}
                    control={control}
                    defaultValue=""
                    render={({onChange, value}) => <TextField
                                                        onChange={onChange}
                                                        value={value}
                                                        disabled={pending}
                                                        required
                                                    />}
                />
                {
                    error.length > 0
                    &&
                    <DialogContentText style={{marginTop: '.5em', marginBottom: '0'}}>{error}</DialogContentText>
                }
            </DialogContent>
            <DialogActions>
                <Button onClick={onClose} disabled={pending}>{i18n.t('browser.create_folder.cancel')}</Button>
                <Button color="primary" type="submit" disabled={pending}>{i18n.t('browser.create_folder.create')}</Button>
            </DialogActions>
        </form>
    </Dialog>;
}