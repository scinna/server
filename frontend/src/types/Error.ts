export type ScinnaError = {
    Message: string;
    ErrCode: number;
    status: number;
}

export function isScinnaError(data: any): boolean {
    return data && data.hasOwnProperty('Message') && data.hasOwnProperty('ErrCode');
}
