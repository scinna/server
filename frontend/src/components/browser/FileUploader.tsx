import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    InputLabel,
    TextareaAutosize,
    TextField
} from "@material-ui/core";
import i18n from "i18n-js";
import React, {useState} from "react";
import {useToken} from "../../context/TokenProvider";
import {Controller, useForm} from "react-hook-form";
import {VisibilityDropDown} from "../VisibilityDropDown";

import styles from "../../assets/scss/browser/Browser.module.scss";

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
    const [pending, setPending] = useState<boolean>(false);
    const {control, handleSubmit} = useForm<IFormInputs>();

    const onSubmit = async (data: IFormInputs) => {
        await setPending(true);
        console.log(data);

        await setPending(false);
    }

    return <Dialog open={shown} onClose={onClose}>
        <form onSubmit={handleSubmit(onSubmit)}>
            <DialogTitle>{i18n.t('browser.file_uploader.title')}</DialogTitle>
            <DialogContent>
                <InputLabel>{i18n.t('browser.file_uploader.file_name')}: </InputLabel>
                <Controller
                    name={"name"}
                    control={control}
                    defaultValue=""
                    render={({onChange, value}) =>
                        <TextField
                            onChange={onChange}
                            value={value}
                            disabled={pending}
                            fullWidth={true}
                            required
                        />}
                />
                <InputLabel
                    className={styles.FileUploader__VisibilityLabel}>{i18n.t('browser.file_uploader.description')}: </InputLabel>
                <Controller
                    name={"description"}
                    control={control}
                    defaultValue={''}
                    render={({onChange, value}) => <TextareaAutosize
                        rowsMin={3}
                        value={value}
                        onChange={onChange}
                    />
                    }
                />
                <InputLabel
                    className={styles.FileUploader__VisibilityLabel}>{i18n.t('browser.file_uploader.visibility')}: </InputLabel>
                <Controller
                    name={"visibility"}
                    control={control}
                    defaultValue={0}
                    render={({onChange, value}) => <VisibilityDropDown
                        selectedVisibility={value}
                        setSelectedVisibility={onChange}
                    />
                    }
                />
            </DialogContent>
            <DialogActions>
                <Button onClick={onClose} disabled={pending}>{i18n.t('browser.file_uploader.cancel')}</Button>
                <Button color="primary" type="submit"
                        disabled={pending}>{i18n.t('browser.file_uploader.upload')}</Button>
            </DialogActions>
        </form>
    </Dialog>
}