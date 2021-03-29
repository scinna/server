import {
    Button,
    Dialog,
    DialogActions,
    DialogContent, DialogContentText,
    DialogTitle, InputLabel,
    TextField
}                                   from "@material-ui/core";
import i18n                         from "i18n-js";
import React, {useState}            from "react";
import {useForm, Controller}        from "react-hook-form";
import {apiCall}                    from "../../utils/useApi";
import {useToken}                   from "../../context/TokenProvider";
import {useBrowser}                 from "../../context/BrowserProvider";
import {isScinnaError, ScinnaError} from "../../types/Error";
import {VisibilityDropDown}         from "../VisibilityDropDown";

import styles       from '../../assets/scss/browser/Browser.module.scss';
import {Collection} from "../../types/Collection";
import {useModal}   from "../../context/ModalProvider";

type Props = {
    closeCallback: () => void;
    collection?: Collection;
}

type IFormInputs = {
    title: string;
    visibility: number;
}

export function FolderEditor({collection, closeCallback = () => {}}: Props) {
    const {hide} = useModal();

    const {username, path, refresh} = useBrowser();
    const {token} = useToken();
    const [pending, setPending] = useState<boolean>(false);
    const [error, setError] = useState<String>("");
    const {control, handleSubmit} = useForm<IFormInputs>();

    const fullyHide = async () => {
        closeCallback();
        await hide();
    }

    const onSubmit = async (data: IFormInputs) => {
        await setPending(true);

        let response;
        if (!collection) {
            response = await apiCall(token, {
                url: '/api/browse/' + username + '/' + path + (path && path?.length > 0 ? '/' : '') + data.title,
                method: 'POST',
                data: {
                    Visibility: data.visibility,
                }
            });
        } else {
            response = await apiCall(token, {
                url: '/api/browse/' + username + '/' + path + (path && path?.length > 0 ? '/' : '') + collection.Title,
                method: 'PUT',
                data: {
                    Title: data.title,
                    Visibility: data.visibility,
                }
            })
        }

        if (isScinnaError(response)) {
            await setError((response as ScinnaError).Message);
            await setPending(false);
            await closeCallback();
            return;
        }

        await setPending(false);
        await setError('');
        await refresh();
        await fullyHide();
    };

    return <Dialog open={true} onClose={fullyHide}>
        <form onSubmit={handleSubmit(onSubmit)}>
            <DialogTitle>{
                collection ?
                    i18n.t('browser.folder_editor.edit_title')
                    :
                    i18n.t('browser.folder_editor.create_title')
            }</DialogTitle>
            <DialogContent>
                <InputLabel>{i18n.t('browser.folder_editor.folder_name')}: </InputLabel>
                <Controller
                    name={"title"}
                    control={control}
                    defaultValue={collection?.Title}
                    render={({onChange, value}) =>
                        <TextField
                            onChange={onChange}
                            value={value}
                            disabled={pending}
                            required
                        />}
                />
                <InputLabel
                    className={styles.CreateFolder__VisibilityLabel}>{i18n.t('browser.folder_editor.visibility')}: </InputLabel>
                <Controller
                    name={"visibility"}
                    control={control}
                    defaultValue={collection ? collection.Visibility : 0}
                    render={({onChange, value}) => <VisibilityDropDown
                        disabled={pending}
                        selectedVisibility={value}
                        setSelectedVisibility={onChange}
                    />
                    }
                />
                {
                    error.length > 0
                    &&
                    <DialogContentText style={{marginTop: '.5em', marginBottom: '0'}}>{error}</DialogContentText>
                }
            </DialogContent>
            <DialogActions>
                <Button onClick={fullyHide} disabled={pending}>{i18n.t('browser.folder_editor.cancel')}</Button>
                <Button color="primary" type="submit"
                        disabled={pending}>{
                    collection ?
                        i18n.t('browser.folder_editor.save')
                        :
                        i18n.t('browser.folder_editor.create')
                }</Button>
            </DialogActions>
        </form>
    </Dialog>;
}