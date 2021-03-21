import React         from 'react';
import {useParams}   from 'react-router-dom';
import {useApiCall}           from "../utils/useApi";
import {Media, MEDIA_PICTURE} from "../types/Media";

import '../types/scss.d.ts';
import styles       from '../assets/scss/ShowPicture.module.scss';
import {Loader}      from "../components/Loader";
import {ShowPicture} from "../components/medias/ShowPicture";

interface ShowPictureParams {
    pictureId: string;
}

const GetComponentForMediaType = (media: Media) => {
    switch(media.MediaType) {
        case MEDIA_PICTURE:
            return <ShowPicture media={media}/>
        default:
            return <span>Unknown media type: {media.MediaType}</span>
    }
}

export function ShowMedia() {
    const {pictureId} = useParams<ShowPictureParams>();
    const info = useApiCall<Media>({url: '/' + pictureId + '/infos', canBeUnauthed: true})

    let clazz = '';
    if (info.status === 'success' && info?.data) {
        clazz = ' ';
        switch (info.data.MediaType) {
            case MEDIA_PICTURE:
                clazz += styles.ShowPicture;
                break;
            default:
                break;
        }
    }

    return <div className={styles.ShowMedia + clazz}>
        {
            info.status === 'success'
            &&
            GetComponentForMediaType(info.data)
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