import {Collection} from "../../types/Collection";
import {Media} from "../../types/Media";
import {Link} from 'react-router-dom';
import {useBrowser} from "../../context/BrowserProvider";

import FolderIcon from '../../assets/images/folder.svg';
import styles from '../../assets/scss/browser/Icon.module.scss';

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
    return <Link className={styles.Icon} to={"/"}>
        <img className={styles.Icon__Image} src={media.Thumbnail} alt=""/>
        <span className={styles.Icon__Text}>{cap(media.Title)}</span>
    </Link>
}

const CollectionIcon = ({collection}: { collection: Collection }) => {
    const {username, path} = useBrowser();
    const correctedPath = (path?.startsWith('/') ? '' : '/') + (path ? path : '');

    return <Link className={styles.Icon} to={"/browse/" + username ?? '' + correctedPath }>
        <img className={styles.Icon__Image} src={FolderIcon} alt={collection.Title} />
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