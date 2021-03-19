import {apiCall} from "../../utils/useApi";
import {isScinnaError} from "../../types/Error";
import {Button, TextField} from "@material-ui/core";
import i18n from "i18n-js";
import {useToken} from "../../context/TokenProvider";
import {useState} from "react";

import styles from '../../assets/scss/server/_InviteGenerator.module.scss';

export function InviteCodeGenerator() {
    const {token} = useToken();
    const [generatedInvite, setGeneratedInvite] = useState<String>("");

    const generate = async () => {
        const response = await apiCall(token, {
            method: 'POST',
            url: '/api/server/admin/invite'
        });

        if (isScinnaError(response)) {
            console.log("Something went wrong");
            return;
        }

        setGeneratedInvite((response as { Code: string }).Code);
    };

    return <div className={styles.InviteGenerator}>
        <TextField disabled={true} label={i18n.t('server_settings.invite.code')} value={generatedInvite}/>
        <Button onClick={generate}>{i18n.t('server_settings.invite.generate')}</Button>
    </div>;
}