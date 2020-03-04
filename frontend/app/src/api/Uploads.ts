import {scinnaxios} from './Axios';

interface IUploadData {
    title: string,
    description: string,
    visibility: number,
    picture?: object|null
}

export interface IUploadResponse {
    Title: string,
    URLID: string,
    Description: string,
    Visibility: string,
    //Creator /** Not used yet? */
}

export function APIUploadPicture(data: IUploadData, updateProgress: Function, actionAfter: Function) { 

    const dataSent = new FormData();
    for (let k in data) {
        // @ts-ignore
        dataSent.append(k, data[k]);
    }

    scinnaxios.post('/pictures', dataSent, {
        onUploadProgress: (evt: any) => {
            updateProgress(evt.loaded / evt.total * 100);
        }
    }).then( (resp) => {
            actionAfter(resp.data)
        })
        .catch((err) => {
            //console.log(err.response.data)
        })
}