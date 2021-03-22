import {Collection} from "../../types/Collection";
import {Media} from "../../types/Media";
import {Link} from 'react-router-dom';
import {useBrowser} from "../../context/BrowserProvider";

export type IconProps = {
    media?: Media;
    collection?: Collection;
};

const MediaIcon = ({media}: { media: Media }) => {
    return <div>
        <img src="" alt=""/>
        <span>MED: {media.Title}</span>
    </div>
}

const CollectionIcon = ({collection}: { collection: Collection }) => {
    const {username, path} = useBrowser();
    return <Link to={"/browse/" + username + "/" + collection.Title}>
        <img src="" alt=""/>
        <span>COL: {collection.Title}</span>
    </Link>
}

export const Icon = ({media, collection}: IconProps) => {
    return <div>
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
    </div>;
}