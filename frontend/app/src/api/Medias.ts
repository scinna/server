import {scinnaxios} from './Axios';

export function APIFetchMediaInfos(media: string|undefined, callback: Function) {
    scinnaxios({ method: 'GET', url: '/medias/' + media})
        .then(resp => {
            callback(resp.data)
        })
        .catch(err => {
            // @TODO: Write a message if the media is not available
            // i.e. the media is private and not owned by the user
        })
}

/** No clue on how to do this for videos */
export function APIFetchPrivateMedia(media: string|undefined, callback: Function) {
    scinnaxios({ method: 'GET', url: '/' + media, responseType: 'arraybuffer'})
        .then(resp => {
            callback(Buffer.from(resp.data, 'binary').toString('base64'))
        })
        .catch(err => {
            // @TODO: Write a message if there is something wrong
        })
}