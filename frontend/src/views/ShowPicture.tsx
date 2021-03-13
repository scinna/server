import React, {useState}            from 'react';
import {useParams}                  from 'react-router-dom';
import {apiCall}                    from "../utils/useApi";
import {Media}                      from "../types/Media";

import '../types/scss.d.ts';
import classes                      from '../assets/scss/ShowPicture.module.scss';
import useAsyncEffect               from "use-async-effect";
import {useToken}                   from "../utils/TokenProvider";
import {isHttpError, isScinnaError} from "../types/Error";
import {Loader}                     from "../components/Loader";

interface ShowPictureParams {
    pictureId: string;
}

export function ShowPicture() {
    const {pictureId} = useParams<ShowPictureParams>();
    const {token} = useToken();
    const [info, setInfo] = useState<Media | null>(null);
    useAsyncEffect(async () => {
        const pictureInfo = await apiCall<Media>(token, {
            url: '/' + pictureId + "/infos",
        })
        if (!pictureInfo) {
            return null;
        }

        if (isScinnaError(pictureInfo)) {
            // @TODO something
            return null;
        } else if (isHttpError(pictureInfo)) {
            // @TODO something
            return null;
        }

        setInfo(pictureInfo as Media);
    }, [pictureId])

    return <div className={classes.ShowPicture}>
        {
            info &&
            <>
                <img className={classes.ShowPicture__Picture} src={"/" + pictureId} alt=""/>
                <div className={classes.ShowPicture__Infos}>
                    <h1>{info.Title}</h1>
                    <p>{info.Description}</p>
                    <span>{info.Author}</span>
                    {
                        info.Visibility !== 2
                        &&
                        <a href={window.location.protocol + '//' + window.location.hostname + '/' + pictureId}>Raw image
                            URL</a>
                    }
                </div>
            </>
        }
        {
            !info
            &&
            <div className="centeredBlock">
                <Loader/>
            </div>
        }

    </div>
}