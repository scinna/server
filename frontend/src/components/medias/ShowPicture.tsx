import React, {useState} from "react";
import {Media}           from "../../types/Media";
import useAsyncEffect from "use-async-effect";
import {useToken}                   from "../../context/TokenProvider";
import {isScinnaError, ScinnaError} from "../../types/Error";
import {Loader}                     from "../Loader";
import {getVisibilityFromNumber} from "../../utils/Mappings";
import {MediaShare}              from "./Share";

import classes                   from "../../assets/scss/ShowPicture.module.scss";

type Props = {
    media: Media;
}

const ShowPrivatePicture = ({media}: Props) => {
    const {token} = useToken();
    const [blobUrl, setBlobUrl] = useState<'pending'|ScinnaError|string>('pending');

    useAsyncEffect(async () => {
        setBlobUrl('pending');
        try {
            const resp = await fetch('/' + media.MediaID, {
                headers: {
                    Authorization: 'Bearer ' + token,
                }
            });

            const img = await resp.blob();
            setBlobUrl(URL.createObjectURL(img));
        } catch (e) {
            console.log(e);
        }
    }, []);

    const isScinnaErr = isScinnaError(blobUrl);

    return <>
        {
            blobUrl === 'pending'
            &&
            <Loader/>
        }
        {
            isScinnaErr
            &&
            <span>{(blobUrl as ScinnaError).Message}</span>
        }
        {
            blobUrl !== 'pending' && !isScinnaErr
            &&
            <img className={classes.ShowPicture__Picture} src={(blobUrl as string)} alt={media.Title}/>
        }
    </>;
}

export function ShowPicture({media}: Props) {
    return <>
        {
            media.Visibility === 2
            &&
            <ShowPrivatePicture media={media}/>
        }
        {
            media.Visibility !== 2
            &&
            <img className={classes.ShowPicture__Picture} src={"/" + media.MediaID} alt=""/>
        }
        <>
            <div className={classes.ShowPicture__Infos}>
                <span className={classes.ShowPicture__Infos__Visibility}>({getVisibilityFromNumber(media.Visibility)})</span>
                <span className={classes.ShowPicture__Infos__Collection}>
                    {'/' + media.Author + '/' + media.Collection}
                </span>
                <h1>{media.Title}</h1>
                <p>{media.Description}</p>
            </div>

            {
                media.Visibility !== 2
                &&
                <MediaShare media={media}/>
            }
        </>
    </>
}