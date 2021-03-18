import styles from '../../assets/scss/ServerSettings.module.scss';
import {useToken} from "../../utils/TokenProvider";
import {Button, TextField} from "@material-ui/core";
import i18n from "i18n-js";
import {apiCall} from "../../utils/useApi";
import {useState} from "react";
import {isScinnaError} from "../../types/Error";

function Generator() {
    const {token} = useToken();
    const [generatedToken, setGeneratedToken] = useState<String>("");

    const generate = async () => {
        const response = await apiCall(token, {
            method: 'POST',
            url: '/api/server/admin/invite'
        });

        if (isScinnaError(response)) {
            console.log("Something went wrong");
            return;
        }

        setGeneratedToken((response as { Code: string }).Code);
    };

    return <div className={styles.InviteCodesTab__Generator}>
        <TextField disabled={true} label={i18n.t('server_settings.invite.code')} value={generatedToken}/>
        <Button onClick={generate}>{i18n.t('server_settings.invite.generate')}</Button>
    </div>;
}

export function InviteCodes() {
    return <div className={styles.Tabbed}>
        <Generator/>

    </div>;
}