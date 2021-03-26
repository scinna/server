import {LinkShortener} from "../components/shortener/LinkShortener";
import Link from "../components/shortener/Link";

import styles from '../assets/scss/linkshortener/shortener.module.scss';
import {useShortenLink} from "../context/ShortenLinkProvider";
import {Loader} from "../components/Loader";

export function UrlShortnener() {
    const {status, links, error} = useShortenLink();

    return <div className={styles.LinkShortener}>
        <LinkShortener/>

        <div className={styles.LinkShortener__List}>
            { status }
            {
                status === 'pending'
                &&
                <Loader/>
            }

            {
                status === 'error'
                &&
                <span>{error}</span>
            }

            {
                status === 'success' && links
                &&
                links.map(l => <Link link={l} />)
            }
        </div>
    </div>;
}