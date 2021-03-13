import React        from 'react';
import {useParams}  from 'react-router-dom';
import {useApiCall} from "../utils/useApi";
import {Media}      from "../types/Media";

import '../types/scss.d.ts';
import classes      from '../assets/scss/ShowPicture.module.scss';
import {Loader}     from "../components/Loader";

interface ShowPictureParams {
    pictureId: string;
}

export function ShowPicture() {
    const {pictureId} = useParams<ShowPictureParams>();
    const info = useApiCall<Media>({url: '/' + pictureId + '/infos'})

    return <div className={classes.ShowPicture}>
        {
            info.status === 'success'
            &&
            <>
                <img className={classes.ShowPicture__Picture} src={"/" + pictureId} alt=""/>
                <div className={classes.ShowPicture__Infos}>
                    <h1>{info.data.Title}</h1>
                    <p>{info.data.Description}</p>
                    <span>{info.data.Author}</span>
                    {
                        info.data.Visibility !== 2
                        &&
                        <a href={window.location.protocol + '//' + window.location.hostname + '/' + pictureId}>
                            Raw image URL
                        </a>
                    }
                </div>
            </>
        }
        {
            info.status === 'pending'
            &&
            <div className="centeredBlock">
                <Loader/>
            </div>
        }
        {
            info.status === 'error'
            &&
            <div>
                {
                    info.error.Message
                }
            </div>
        }

    </div>
}