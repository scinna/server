import React from 'react';
import {useParams} from "react-router-dom";
import {useApiCall} from "../utils/useApi";
import {Loader} from "../components/Loader";

import styles from '../assets/scss/validate/index.module.scss';
import i18n from "i18n-js";

type Params = {
    valCode: string;
}

type ValidationResponse = {
    username: string;
}

export const ValidateAccount = () => {
    const {valCode} = useParams<Params>();

    const apiResponse = useApiCall<ValidationResponse>({
        url: '/api/auth/register/' + valCode,
        canBeUnauthed: true,
    })

    return <div className={styles.Validate}>
        {
            apiResponse.status === 'pending'
            &&
            <Loader/>
        }

        {
            apiResponse.status === 'success'
            &&
            <>
                <p>{i18n.t('validate.accountValidated')}.</p><p>{i18n.t('validate.connect')} <span>{apiResponse.data.username}</span>.</p>
            </>
        }

        {
            apiResponse.status === 'error'
            &&
            <p>
                {apiResponse.error.Message}
            </p>
        }
    </div>
}