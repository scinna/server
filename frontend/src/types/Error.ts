export type ScinnaError = {
    Message: string;
    ErrCode: number;
}

export type HttpError = {
    status: number;
}

export function isScinnaError(data: object): boolean {
    return data.hasOwnProperty('Message') && data.hasOwnProperty('ErrCode');
}

export function isHttpError(data: object): boolean {
    return data.hasOwnProperty("status");
}