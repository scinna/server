export const MEDIA_PICTURE = 0;
export const MEDIA_VIDEO = 1;
export const MEDIA_TXTBIN = 2;
export const MEDIA_SHORTURL = 3;

export type Media = {
    MediaID: string;
    MediaType: number;
    Title: string;
    Description: string;
    Visibility: number;
    PublishedAt: Date;
    Collection: string;
    Author: string;
};