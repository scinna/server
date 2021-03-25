import {IconButton, TextField, TextFieldProps} from "@material-ui/core";
import {FileCopy} from '@material-ui/icons';
import {useRef} from "react";

import styles from '../assets/scss/components/copiable_textfield.module.scss'

export function CopiableTextfield({...args}: TextFieldProps) {
    const textField = useRef();

    const copyToClipboard = () => {
        // not in my best mood so idgas about this function.
        // @todo maybe fix one day

        //@ts-ignore
        textField?.current?.select();
        //@ts-ignore
        textField?.current?.setSelectionRange(0, 999999);
        document.execCommand('copy');
    }

    return <div className={styles.CopiableTextfield}>
        <TextField contentEditable={false} {...args} inputRef={textField} />
        <IconButton className={styles.CopiableTextfield__Button} color="primary" onClick={copyToClipboard}>
            <FileCopy />
        </IconButton>
    </div>
}