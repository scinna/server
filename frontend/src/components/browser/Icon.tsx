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
import {useIconContextMenu}         from "../../context/IconContextMenuProvider";

import PrivateIcon  from '../../assets/images/private.svg';
import UnlistedIcon from '../../assets/images/unlisted.svg';
import PublicIcon   from '../../assets/images/public.svg';

import styles                    from '../../assets/scss/browser/Icon.module.scss';
import {getVisibilityFromNumber} from "../../utils/Mappings";

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

const VisibilityBadge = ({visibility}: { visibility: number }) => {
    const getIcon = () => {
        switch (visibility) {
            case 2:
                return PrivateIcon;
            case 1:
                return UnlistedIcon;
            default:
                return PublicIcon;
        }
    }

    return <img className={styles.Icon__Badge} src={getIcon()} alt={getVisibilityFromNumber(visibility)}/>;
}

const MediaIcon = ({media}: { media: Media }) => {
    const {show} = useIconContextMenu();
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
        <VisibilityBadge visibility={media.Visibility}/>
    </Link>
}

const CollectionIcon = ({collection}: { collection: Collection }) => {
    const {show} = useIconContextMenu();
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
        <VisibilityBadge visibility={collection.Visibility}/>
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