import {Media} from '@/types/Media';

export type Collection = {
    CollectionID: string;
    Title: string;
    Visibility: number;
    Collections: Collection[];
    Medias: Media[];
}