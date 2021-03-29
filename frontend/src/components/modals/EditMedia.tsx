import {useModal} from "../../context/ModalProvider";
import {useToken} from "../../context/TokenProvider";
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    InputLabel, TextareaAutosize,
    TextField
}                                   from "@material-ui/core";
import i18n                         from "i18n-js";
import {Controller, useForm}        from "react-hook-form";
import styles                       from "../../assets/scss/browser/Browser.module.scss";
import {VisibilityDropDown}         from "../VisibilityDropDown";
import React, {useState}            from "react";
import {Media}                      from "../../types/Media";
import {apiCall}                    from "../../utils/useApi";
import {isScinnaError, ScinnaError} from "../../types/Error";
import {useBrowser}                 from "../../context/BrowserProvider";

type IFormInputs = {
    title: string;
    description: string;
    visibility: number;
};

type Props = {
  media: Media;
  closeCallback: () => void;
};

export function EditMedia({media, closeCallback = () => {}}: Props) {
    const {hide} = useModal();
    const {token} = useToken();
    const {refresh} = useBrowser();
    const [pending, setPending] = useState<boolean>(false);
    const [error, setError] = useState<String>("");
    const {control, handleSubmit} = useForm<IFormInputs>();

    const close = async () => {
        closeCallback();
        await hide();
    }

    const onSubmit = async (data: IFormInputs) => {
        await setPending(true);
        await setError('');

        let fd = new FormData();
        Object.entries(data).forEach(([key, val]) => {
            fd.append(key, '' + val);
        });

        const response = await apiCall(token, {
           url: '/' + media.MediaID,
            method: 'PUT',
            data: fd,
            mustNotBeStringified: true,
        });

        if (isScinnaError(response)) {
            await setPending(false);
            await setError((response as ScinnaError).Message);

            return;
        }

        await refresh();
        await close();
    }

    return <Dialog open={true} onClose={close}>
        <form onSubmit={handleSubmit(onSubmit)}>
            <DialogTitle>{i18n.t('browser.modals.edit_media.title')}</DialogTitle>
            <DialogContent>
                <InputLabel>{i18n.t('browser.modals.edit_media.media_title')}: </InputLabel>
                <Controller
                    name={"title"}
                    control={control}
                    defaultValue={media.Title}
                    render={({onChange, value}) =>
                        <TextField
                            onChange={onChange}
                            value={value}
                            disabled={pending}
                            required
                        />}
                />
                <InputLabel className={styles.FileUploader__VisibilityLabel}>
                    {i18n.t('browser.modals.edit_media.description')}:
                </InputLabel>
                <Controller
                    name={"description"}
                    control={control}
                    defaultValue={media.Description}
                    render={({onChange, value}) => <TextareaAutosize
                        className={styles.FileUploader__Description}
                        disabled={pending}
                        rowsMin={3}
                        value={value}
                        onChange={onChange}
                    />
                    }
                />
                <InputLabel
                    className={styles.CreateFolder__VisibilityLabel}>{i18n.t('browser.modals.edit_media.visibility')}: </InputLabel>
                <Controller
                    name={"visibility"}
                    control={control}
                    defaultValue={media.Visibility}
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
                <Button onClick={close} disabled={pending}>{i18n.t('browser.folder_editor.cancel')}</Button>
                <Button color="primary" type="submit"disabled={pending}>{i18n.t('browser.folder_editor.save')}</Button>
            </DialogActions>
        </form>
    </Dialog>;
}