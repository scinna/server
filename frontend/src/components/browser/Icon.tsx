import {useState}                   from "react";
import {Collection}                 from "../../types/Collection";
import {Media}                      from "../../types/Media";
import {Link}                       from 'react-router-dom';
import {useBrowser}                 from "../../context/BrowserProvider";
import FolderIcon                   from '../../assets/images/folder.svg';
import InaccessibleIcon             from '../../assets/images/inaccessible_icon.png';
import {useToken}                   from "../../context/TokenProvider";
import {isScinnaError, ScinnaError} from "../../types/Error";
import useAsyncEffect               from "use-async-effect";

import styles           from '../../assets/scss/browser/Icon.module.scss';
import {useIconContext} from "../../context/IconContextProvider";

export type IconProps = {
    media?: Media;
    collection?: Collection;
};

const cap = (title: string): string => {
    if (title.length > 27) {
        return title.substr(0, 24) + "...";
    }

    return title;
}

const MediaIcon = ({media}: { media: Media }) => {
    const {show} = useIconContext();
    const thumbnailRawUrl = "/" + media.MediaID + "/thumbnail";
    const {token} = useToken();
    const [thumbnailUrl, setThumbnailUrl] = useState<'pending' | ScinnaError | string>('pending');

    useAsyncEffect(async () => {
        if (media.Visibility !== 2) {
            await setThumbnailUrl(thumbnailRawUrl);
            return;
        }

        setThumbnailUrl('pending');
        try {
            const resp = await fetch(thumbnailRawUrl, {
                headers: {
                    Authorization: 'Bearer ' + token,
                }
            });

            const img = await resp.blob();
            setThumbnailUrl(URL.createObjectURL(img));
        } catch (e) {
            console.log(e);
        }
    }, []);

    const isErr = isScinnaError(thumbnailUrl);

    return <Link className={styles.Icon} to={"/"} onContextMenu={show(null, media)}>
        {
            isErr
            &&
            <img className={styles.Icon__Image} src={InaccessibleIcon} alt=""/>
        }
        {
            !isErr
            &&
            <img className={styles.Icon__Image} src={thumbnailUrl as string} alt=""/>
        }
        <span className={styles.Icon__Text}>{cap(media.Title)}</span>
    </Link>
}

const CollectionIcon = ({collection}: { collection: Collection }) => {
    const {show} = useIconContext();
    const {username, path} = useBrowser();
    let fullPath = "/browse/" + (username ?? '') + '/';

    if (path?.length === 0) {
        fullPath += collection.Title;
    } else {
        fullPath += path + (!path?.endsWith('/') ? '/' : '') + collection.Title;
    }

    return <Link className={styles.Icon} to={fullPath} onContextMenu={show(collection, null)}>
        <img className={styles.Icon__Image} src={FolderIcon} alt={collection.Title}/>
        <span className={styles.Icon__Text}>{cap(collection.Title)}</span>
    </Link>
}

export const Icon = ({media, collection}: IconProps) => {
    return <>
        {
            media
            &&
            <MediaIcon media={media}/>
        }
        {
            collection
            &&
            <CollectionIcon collection={collection}/>
        }
    </>;
}